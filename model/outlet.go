package model

type Outlet interface {
	Switch

	SetInUse(bool)
	IsInUse() bool

	// Sets the on state changed callback
	InUseStateChanged(func(bool))
}
