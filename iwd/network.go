package iwd

import dbus "github.com/godbus/dbus/v5"

// Networks refers to the output of iwd.networks.Station.GetOrderedNetworks
type Networks []struct {
	dbus.ObjectPath
	Type int
}

// Network refers to the iwd network for example: /net/connman/iwd/0/4/7a65696b7561697a65696b756169646577616e67_psk
type Network struct {
	Connected    bool
	Device       dbus.ObjectPath
	KnownNetwork dbus.ObjectPath
	Name         string
	Type         string
}
