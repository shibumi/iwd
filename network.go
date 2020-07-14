package go-iwd

import dbus "github.com/godbus/dbus/v5"

const (
	objectNetwork = "net.connman.iwd.Network"
)

// Network refers to the iwd network for example: /net/connman/iwd/0/4/7a65696b7561697a65696b756169646577616e67_psk
type Network struct {
	Connected    bool
	Device       dbus.ObjectPath
	KnownNetwork dbus.ObjectPath
	Name         string
	Type         string
}
