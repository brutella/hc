package dnssd

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd/log"
	"github.com/miekg/dns"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"net"
)

var (
	IPv4LinkLocalMulticast = net.ParseIP("224.0.0.251")
	IPv6LinkLocalMulticast = net.ParseIP("ff02::fb")

	AddrIPv4LinkLocalMulticast = &net.UDPAddr{
		IP:   IPv4LinkLocalMulticast,
		Port: 5353,
	}

	AddrIPv6LinkLocalMulticast = &net.UDPAddr{
		IP:   IPv6LinkLocalMulticast,
		Port: 5353,
	}

	TtlDefault  uint32 = 75 * 60 // Default ttl for mDNS resource records
	TtlHostname uint32 = 120     // TTL for mDNS resource records containing the host name
)

// Query is a mDNS query
type Query struct {
	msg *dns.Msg // The query message
}

// Response is a mDNS response
type Response struct {
	msg  *dns.Msg     // The response message
	addr *net.UDPAddr // Is nil for multicast response
}

// Request represents an incoming mDNS message
type Request struct {
	msg   *dns.Msg       // The message
	from  *net.UDPAddr   // The source addr of the message
	iface *net.Interface // The network interface from which the message was received
}

// MDNSConn represents a mDNS connection. It encapsulates an IPv4 and IPv6 UDP connection.
type MDNSConn interface {
	// SendQuery sends a mDNS query.
	SendQuery(q *Query) error

	// SendResponse sends a mDNS response
	SendResponse(resp *Response) error

	// Read returns a channel which receives mDNS messages
	Read(ctx context.Context) <-chan *Request

	// Close closes the connection
	Close()
}

type mdnsConn struct {
	ipv4 *ipv4.PacketConn
	ipv6 *ipv6.PacketConn
	ch   chan *Request
}

func NewMDNSConn() (MDNSConn, error) {
	return newMDNSConn()
}

func (c *mdnsConn) SendQuery(q *Query) error {
	return c.sendQuery(q.msg)
}

func (c *mdnsConn) SendResponse(resp *Response) error {
	if resp.addr != nil {
		return c.sendResponseTo(resp.msg, resp.addr)
	}

	return c.sendResponse(resp.msg)
}

func (c *mdnsConn) Read(ctx context.Context) <-chan *Request {
	return c.read(ctx)
}

func (c *mdnsConn) Close() {
	c.close()
}

func newMDNSConn() (*mdnsConn, error) {
	var errs []error
	var connIPv4 *ipv4.PacketConn
	var connIPv6 *ipv6.PacketConn

	if conn, err := net.ListenUDP("udp4", AddrIPv4LinkLocalMulticast); err != nil {
		errs = append(errs, err)
	} else {
		connIPv4 = ipv4.NewPacketConn(conn)
		connIPv4.SetControlMessage(ipv4.FlagInterface, true)

		// Don't send us our own messages back
		connIPv4.SetMulticastLoopback(false)
	}

	if conn, err := net.ListenUDP("udp6", AddrIPv6LinkLocalMulticast); err != nil {
		errs = append(errs, err)
	} else {
		connIPv6 = ipv6.NewPacketConn(conn)
		connIPv6.SetControlMessage(ipv6.FlagInterface, true)

		// Don't send us our own messages back
		connIPv6.SetMulticastLoopback(false)
	}

	if err := first(errs...); connIPv4 == nil && connIPv6 == nil {
		return nil, fmt.Errorf("Failed setting up UDP server: %v", err)
	}

	return &mdnsConn{
		ipv4: connIPv4,
		ipv6: connIPv6,
		ch:   make(chan *Request),
	}, nil
}

func (c *mdnsConn) close() {
	if c.ipv4 != nil {
		c.ipv4.Close()
	}

	if c.ipv6 != nil {
		c.ipv6.Close()
	}
}

func (c *mdnsConn) read(ctx context.Context) <-chan *Request {
	c.readInto(ctx, c.ch)
	return c.ch
}

func (c *mdnsConn) readInto(ctx context.Context, ch chan *Request) {
	var isReading = true
	if c.ipv4 != nil {
		go func() {
			buf := make([]byte, 65536)
			for {
				if !isReading {
					return
				}

				n, cm, from, err := c.ipv4.ReadFrom(buf)
				if err != nil {
					continue
				}

				udpAddr, ok := from.(*net.UDPAddr)
				if !ok {
					log.Info.Println("invalid source address")
					continue
				}

				iface, err := net.InterfaceByIndex(cm.IfIndex)
				if err != nil {
					continue
				}

				if n > 0 {
					m := new(dns.Msg)
					if err := m.Unpack(buf); err == nil && !shouldIgnore(m) {
						ch <- &Request{m, udpAddr, iface}
					}
				}
			}
		}()
	}

	if c.ipv6 != nil {
		go func() {
			buf := make([]byte, 65536)
			for {
				if !isReading {
					return
				}

				n, cm, from, err := c.ipv6.ReadFrom(buf)
				if err != nil {
					continue
				}

				udpAddr, ok := from.(*net.UDPAddr)
				if !ok {
					log.Info.Println("invalid source address")
					continue
				}

				iface, err := net.InterfaceByIndex(cm.IfIndex)
				if err != nil {
					continue
				}

				if n > 0 {
					m := new(dns.Msg)
					if err := m.Unpack(buf); err == nil && !shouldIgnore(m) {
						ch <- &Request{m, udpAddr, iface}
					}
				}
			}
		}()
	}

	go func() {
		<-ctx.Done()
		isReading = false
	}()
}

func (c *mdnsConn) sendQuery(m *dns.Msg) error {
	sanitizeQuery(m)

	return c.writeMsg(m)
}

func (c *mdnsConn) sendResponse(m *dns.Msg) error {
	sanitizeResponse(m)

	return c.writeMsg(m)
}

func (c *mdnsConn) sendResponseTo(m *dns.Msg, addr *net.UDPAddr) error {
	sanitizeResponse(m)

	return c.writeMsgTo(m, addr)
}

func (c *mdnsConn) writeMsg(m *dns.Msg) error {
	var err error
	if c.ipv4 != nil {
		err = c.writeMsgTo(m, AddrIPv4LinkLocalMulticast)
	}

	if c.ipv6 != nil {
		err = c.writeMsgTo(m, AddrIPv6LinkLocalMulticast)
	}

	return err
}

func (c *mdnsConn) writeMsgTo(m *dns.Msg, addr *net.UDPAddr) error {
	sanitizeMsg(m)

	var err error
	if c.ipv4 != nil && addr.IP.To4() != nil {
		if out, err := m.Pack(); err == nil {
			_, err = c.ipv4.WriteTo(out, nil, addr)
		}
	}

	if c.ipv6 != nil && addr.IP.To4() == nil {
		if out, err := m.Pack(); err == nil {
			_, err = c.ipv6.WriteTo(out, nil, addr)
		}
	}

	return err
}

func shouldIgnore(m *dns.Msg) bool {
	if m.Opcode != 0 {
		return true
	}

	if m.Rcode != 0 {
		return true
	}

	return false
}

func sanitizeResponse(m *dns.Msg) {
	if m.Question != nil && len(m.Question) > 0 {
		log.Info.Println("Multicast DNS responses MUST NOT contain any questions in the Question Section.  (RFC6762 6)")
		m.Question = nil
	}

	if !m.Response {
		log.Info.Println("In response messages the QR bit MUST be one (RFC6762 18.2)")
		m.Response = true
	}

	if !m.Authoritative {
		log.Info.Println("AA Bit bit MUST be set to one in response messages (RFC6762 18.4)")
		m.Authoritative = true
	}

	if m.Truncated {
		log.Info.Println("In multicast response messages, the TC bit MUST be zero on transmission. (RFC6762 18.5)")
		m.Truncated = false
	}
}

func sanitizeQuery(m *dns.Msg) {
	if m.Response {
		log.Info.Println("In query messages the QR bit MUST be zero (RFC6762 18.2)")
		m.Response = false
	}

	if m.Authoritative {
		log.Info.Println("AA Bit MUST be zero in query messages (RFC6762 18.4)")
		m.Authoritative = false
	}
}

func sanitizeMsg(m *dns.Msg) {
	if m.Opcode != 0 {
		log.Info.Println("In both multicast query and multicast response messages, the OPCODE MUST be zero on transmission (RFC6762 18.3)")
		m.Opcode = 0
	}

	if m.RecursionDesired {
		log.Info.Println("In both multicast query and multicast response messages, the Recursion Available bit MUST be zero on transmission. (RFC6762 18.7)")
		m.RecursionDesired = false
	}

	if m.Zero {
		log.Info.Println("In both query and response messages, the Zero bit MUST be zero on transmission (RFC6762 18.8)")
		m.Zero = false
	}

	if m.AuthenticatedData {
		log.Info.Println("In both multicast query and multicast response messages, the Authentic Data bit MUST be zero on transmission (RFC6762 18.9)")
		m.AuthenticatedData = false
	}

	if m.CheckingDisabled {
		log.Info.Println("In both multicast query and multicast response messages, the Checking Disabled bit MUST be zero on transmission (RFC6762 18.10)")
		m.CheckingDisabled = false
	}

	if m.Rcode != 0 {
		log.Info.Println("In both multicast query and multicast response messages, the Response Code MUST be zero on transmission. (RFC6762 18.11)")
		m.Rcode = 0
	}
}

func first(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

// Sets the Top Bit of rrclass for all answer records (except PTR) to trigger a cache flush in the receivers.
func setAnswerCacheFlushBit(msg *dns.Msg) {
	// From RFC6762
	//    The most significant bit of the rrclass for a record in the Answer
	//    Section of a response message is the Multicast DNS cache-flush bit
	//    and is discussed in more detail below in Section 10.2, "Announcements
	//    to Flush Outdated Cache Entries".
	for _, a := range msg.Answer {
		switch a.(type) {
		case *dns.PTR:
			continue
		default:
			a.Header().Class |= (1 << 15)
		}
	}
}

// Sets the Top Bit of class to indicate the unicast responses are preferred for this question.
func setQuestionUnicast(q *dns.Question) {
	q.Qclass |= (1 << 15)
}

// Returns true if q requires unicast responses.
func isUnicastQuestion(q dns.Question) bool {
	// From RFC6762
	// 18.12.  Repurposing of Top Bit of qclass in Question Section
	//
	//    In the Question Section of a Multicast DNS query, the top bit of the
	//    qclass field is used to indicate that unicast responses are preferred
	//    for this particular question.  (See Section 5.4.)
	return q.Qclass&(1<<15) != 0
}
