package iwd

import "github.com/godbus/dbus/v5"

const (
	objectAp   = "net.connman.iwd.Ap"
	callApActivate = "net.connman.iwd.Station.Scan"
)

// Station refers to net.connman.iwd.Station
type Ap struct {
	Path             dbus.ObjectPath
	ConnectedNetwork dbus.ObjectPath
	Scanning         bool
	State            string
}