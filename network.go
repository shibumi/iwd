package iwd

import (
	dbus "github.com/godbus/dbus/v5"
	"strconv"
)

const (
	objectNetwork      = "net.connman.iwd.Network"
	callNetworkConnect = "net.connman.iwd.Network.Connect"
)

// Network refers to the iwd network for example: /net/connman/iwd/0/4/7a65696b7561697a65696b756169646577616e67_psk
type Network struct {
	Path         dbus.ObjectPath
	Connected    bool
	Device       dbus.ObjectPath
	KnownNetwork dbus.ObjectPath
	Name         string
	Type         string
}

// Connect establishes a connection with a network
// Currently this only works for open networks
func (n *Network) Connect(conn *dbus.Conn) error {
	// path = /net/connman/iwd/<adapter>/<device>/<hex encoded ssid>_<network type>
	device := conn.Object(objectIwd, n.Path)
	call := device.Call(callNetworkConnect, 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

func GetNetwork(conn *dbus.Conn, objectPath dbus.ObjectPath) (*Network, error) {
	obj := conn.Object(objectNetwork, objectPath)
	name, err := obj.GetProperty("Name")
	if err != nil {
		return nil, err
	}
	networkType, err := obj.GetProperty("Type")
	if err != nil {
		return nil, err
	}
	connected, err := obj.GetProperty("Connected")
	if err != nil {
		return nil, err
	}
	connectedBool, err := strconv.ParseBool(connected.String())
	if err != nil {
		return nil, err
	}
	return &Network{
		Path:         objectPath,
		Connected:    connectedBool,
		Device:       "",
		KnownNetwork: "",
		Name:         name.String(),
		Type:         networkType.String(),
	}, nil
}
