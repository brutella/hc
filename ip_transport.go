package hc

import (
	"bytes"
	"context"
	"image"
	"io/ioutil"
	"net"
	"strings"
	"sync"

	"github.com/brutella/dnssd"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/endpoint"
	"github.com/brutella/hc/hap/http"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/util"
	"github.com/xiam/to"
)

type ipTransport struct {
	CameraSnapshotReq func(width, height uint) (*image.Image, error)

	config  *Config
	context hap.Context
	server  *http.Server
	mutex   *sync.Mutex

	storage  util.Storage
	database db.Database

	device    hap.SecuredDevice
	container *accessory.Container

	// Used to communicate between different parts of the program (e.g. successful pairing with HomeKit)
	emitter event.Emitter

	ctx    context.Context
	cancel context.CancelFunc

	responder dnssd.Responder
	handle    dnssd.ServiceHandle

	stopped chan struct{}
}

// NewIPTransport creates a transport to provide accessories over IP.
//
// The IP transports stores the crypto keys inside a database, which
// is by default inside a folder at the current working directory.
// The folder is named exactly as the accessory name.
//
// The transports can contain more than one accessory. If this is the
// case, the first accessory acts as the HomeKit bridge.
//
// *Important:* Changing the name of the accessory, or letting multiple
// transports store the data inside the same database lead to
// unexpected behavior – don't do that.
//
// The transport is secured with an 8-digit pin, which must be entered
// by an iOS client to successfully pair with the accessory. If the
// provided transport config does not specify any pin, 00102003 is used.
func NewIPTransport(config Config, a *accessory.Accessory, as ...*accessory.Accessory) (*ipTransport, error) {
	// Find transport name which is visible in mDNS
	name := a.Info.Name.GetValue()
	if len(name) == 0 {
		log.Info.Panic("Invalid empty name for first accessory")
	}

	cfg := defaultConfig(name)
	cfg.merge(config)

	var storage util.Storage
	if config.Storage == nil {
		var err error
		storage, err = util.NewFileStorage(cfg.StoragePath)
		if err != nil {
			return nil, err
		}
	} else {
		storage = config.Storage
	}

	database := db.NewDatabaseWithStorage(storage)

	hap_pin, err := NewPin(cfg.Pin)
	if err != nil {
		return nil, err
	}

	cfg.load(storage)

	device, err := hap.NewSecuredDevice(cfg.id, hap_pin, database)
	if err != nil {
		return nil, err
	}

	responder, err := dnssd.NewResponder()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	t := &ipTransport{
		storage:   storage,
		database:  database,
		device:    device,
		config:    cfg,
		container: accessory.NewContainer(),
		mutex:     &sync.Mutex{},
		context:   hap.NewContextForSecuredDevice(device),
		emitter:   event.NewEmitter(),
		responder: responder,
		ctx:       ctx,
		cancel:    cancel,
		stopped:   make(chan struct{}),
	}

	t.addAccessory(a)
	for _, a := range as {
		t.addAccessory(a)
	}

	// Users can only pair discoverable accessories
	if t.isPaired() {
		cfg.discoverable = false
	}

	cfg.categoryId = uint8(t.container.AccessoryType())
	cfg.updateConfigHash(t.container.ContentHash())
	cfg.save(storage)

	// Listen for events to update mDNS txt records
	t.emitter.AddListener(t)

	return t, err
}

func (t *ipTransport) Start() {
	// Create server which handles incoming tcp connections
	config := http.Config{
		Port:      t.config.Port,
		Context:   t.context,
		Database:  t.database,
		Container: t.container,
		Device:    t.device,
		Mutex:     t.mutex,
		Emitter:   t.emitter,
	}

	s := http.NewServer(config)
	t.server = s

	if t.CameraSnapshotReq != nil {
		t.server.Mux.Handle("/resource", endpoint.NewResource(t.context, t.CameraSnapshotReq))
	}

	// Publish server port which might be different then `t.config.Port`
	t.config.servePort = int(to.Int64(s.Port()))

	service := newService(t.config)
	t.handle, _ = t.responder.Add(service)

	mdnsCtx, mdnsCancel := context.WithCancel(t.ctx)
	defer mdnsCancel()

	mdnsStop := make(chan struct{})
	go func() {
		t.responder.Respond(mdnsCtx)
		log.Debug.Println("mdns responder stopped")
		mdnsStop <- struct{}{}
	}()

	// keepAliveCtx, keepAliveCancel := context.WithCancel(t.ctx)
	// defer keepAliveCancel()
	//
	// // Send keep alive notifications to all connected clients every 10 minutes
	// keepAlive := hap.NewKeepAlive(10*time.Minute, t.context)
	// go func() {
	// 	keepAlive.Start(keepAliveCtx)
	// 	log.Info.Println("Keep alive stopped")
	// }()

	// Publish accessory ip
	log.Info.Printf("Listening on port %s\n", s.Port())

	serverCtx, serverCancel := context.WithCancel(t.ctx)
	defer serverCancel()
	serverStop := make(chan struct{})
	go func() {
		s.ListenAndServe(serverCtx)
		log.Debug.Println("server stopped")
		serverStop <- struct{}{}
	}()

	// Wait until mdns responder and server stopped
	<-mdnsStop
	<-serverStop
	t.stopped <- struct{}{}
}

// Stop stops the ip transport by stopping the http server and unpublishing the mDNS service.
func (t *ipTransport) Stop() <-chan struct{} {
	t.cancel()

	return t.stopped
}

// XHMURI returns a X-HM styled uri to easily add the accessory to HomeKit.
// To print a QR code to the console, use the follow code snippet.
// ```
// import "github.com/mdp/qrterminal/v3"
//
// uri, _ := transport.XHMURI()
// qrterminal.Generate(uri, qrterminal.L, os.Stdout)
// ```
func (t *ipTransport) XHMURI() (string, error) {
	return t.config.XHMURI(util.SetupFlagIP)
}

// isPaired returns true when the transport is already paired
func (t *ipTransport) isPaired() bool {

	// If more than one entity is stored in the database, we are paired with a device.
	// The transport itself is a device and is stored in the database, therefore
	// we have to check for more than one entity.
	if es, err := t.database.Entities(); err == nil && len(es) > 1 {
		return true
	}

	return false
}

func (t *ipTransport) updateMDNSReachability() {
	t.config.discoverable = t.isPaired() == false
	if t.handle != nil {
		t.handle.UpdateText(t.config.txtRecords(), t.responder)
	}
}

func (t *ipTransport) addAccessory(a *accessory.Accessory) {
	t.container.AddAccessory(a)

	for _, s := range a.Services {
		for _, c := range s.Characteristics {
			// When a characteristic value changes and events are enabled for this characteristic
			// all listeners are notified. Since we don't track which client is interested in
			// which characteristic change event, we send them to all active connections.
			onConnChange := func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
				if c.Events == true {
					t.notifyListener(a, c, conn)
				}
			}
			c.OnValueUpdateFromConn(onConnChange)

			onChange := func(c *characteristic.Characteristic, new, old interface{}) {
				if c.Events == true {
					t.notifyListener(a, c, nil)
				}
			}
			c.OnValueUpdate(onChange)
		}
	}
}

func (t *ipTransport) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic, except net.Conn) {
	conns := t.context.ActiveConnections()
	for _, conn := range conns {
		if conn == except {
			continue
		}
		resp, err := hap.NewCharacteristicNotification(a, c)
		if err != nil {
			log.Info.Panic(err)
		}

		// Write response into buffer to replace HTTP protocol
		// specifier with EVENT as required by HAP
		var buffer = new(bytes.Buffer)
		resp.Write(buffer)
		bytes, err := ioutil.ReadAll(buffer)
		bytes = hap.FixProtocolSpecifier(bytes)
		log.Debug.Printf("%s <- %s", conn.RemoteAddr(), string(bytes))
		conn.Write(bytes)
	}
}

// Handles event which are sent when pairing with a device is added or removed
func (t *ipTransport) Handle(ev interface{}) {
	switch ev.(type) {
	case event.DevicePaired:
		log.Debug.Printf("Event: paired with device")
		t.updateMDNSReachability()
	case event.DeviceUnpaired:
		log.Debug.Printf("Event: unpaired with device")
		t.updateMDNSReachability()
	default:
		break
	}
}

func newService(config *Config) dnssd.Service {
	// 2016-03-14(brutella): Replace whitespaces (" ") from service name
	// with underscores ("_")to fix invalid http host header field value
	// produces by iOS.
	//
	// [Radar] http://openradar.appspot.com/radar?id=4931940373233664
	stripped := strings.Replace(config.name, " ", "_", -1)

	var ips []net.IP
	if ip := net.ParseIP(config.IP); ip != nil {
		ips = append(ips, ip)
	}

	dnsCfg := dnssd.Config{
		Name:   util.RemoveAccentsFromString(stripped),
		Type:   "_hap._tcp",
		Domain: "local",
		Host:   strings.Replace(config.id, ":", "", -1), // use the id (without the colons) to get unique hostnames
		Text:   config.txtRecords(),
		IPs:    ips,
		Port:   config.servePort,
	}
	service, err := dnssd.NewService(dnsCfg)
	if err != nil {
		log.Info.Fatal(err)
	}

	return service
}
