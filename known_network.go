package iwd

const (
	objectKnownNetwork = "net.connman.iwd.KnownNetwork"
)

// KnownNetwork refers to the net.connman.iwd.KnownNetwork object
type KnownNetwork struct {
	AutoConnect       bool
	Hidden            bool
	LastConnectedTime string
	Name              string
	Type              string
}
