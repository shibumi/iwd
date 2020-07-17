package iwd

import "github.com/godbus/dbus/v5"

const (
	objectKnownNetwork = "net.connman.iwd.KnownNetwork"
)

// KnownNetwork refers to the net.connman.iwd.KnownNetwork object
type KnownNetwork struct {
	Path              dbus.ObjectPath
	AutoConnect       bool
	Hidden            bool
	LastConnectedTime string
	Name              string
	Type              string
}
