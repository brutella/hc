package model

// An outlet is a switch which additionally has an inUse state
type Outlet interface {
	Switch

	// SetInUse sets the inUse state
	SetInUse(bool)

	// IsInUse returns true when the outlet is in use, otherwise false
	IsInUse() bool

	// InUseStateChanged calls the argument function when the outlet's
	// isUse state changed
	InUseStateChanged(func(bool))
}
