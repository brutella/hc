package db

type Dns interface {
	Name() string
	Configuration() int64
	SetConfiguration(int64)

	State() int64
	SetState(int64)
}

type dns struct {
	name          string
	configuration int64
	state         int64
}

func NewDns(name string, configuration, state int64) *dns {
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
