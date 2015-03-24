package db

// DNS contains the mDNS TXT records of a HomeKit bridge.
type DNS interface {
	// Name returns the bridge name
	Name() string

	// Configuration returns the bridge configuration (appears as #c in TXT record)
	Configuration() int64

	// SetConfiguration sets the bridge configuration
	SetConfiguration(int64)

	// State returns the bridge state (appears as #s in TXT record)
	State() int64

	// SetState sets the bridge state
	SetState(int64)
}

type dns struct {
	name          string
	configuration int64
	state         int64
}

// NewDNS returns a dns with name, configuration and state.
func NewDNS(name string, configuration, state int64) DNS {
	return &dns{name, configuration, state}
}

func (d *dns) Name() string {
	return d.name
}

func (d *dns) Configuration() int64 {
	return d.configuration
}

func (d *dns) SetConfiguration(c int64) {
	d.configuration = c
}

func (d *dns) State() int64 {
	return d.state
}

func (d *dns) SetState(s int64) {
	d.state = s
}
