package accessory

type Bridge struct {
	*Accessory
}

// NewBridge returns a bridge which implements model.Bridge.
func NewBridge(info Info) *Bridge {
	acc := Bridge{}
	acc.Accessory = New(info, TypeBridge)

	return &acc
}
