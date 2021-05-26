package iwd

import "github.com/godbus/dbus/v5"

const (
	objectAp       = "net.connman.iwd.AccessPoint"
	callApActivate = "net.connman.iwd.Station.Start"
)

// Station refers to net.connman.iwd.Station
type Ap struct {
	Path             dbus.ObjectPath
	ConnectedNetwork dbus.ObjectPath
	State            string
}

func (a Ap) Activateconn(conn *dbus.Conn) error {
	obj := conn.Object(objectAp, "")
	call := obj.Call(callStationScan, 0, "5g", "afoe11afoe11")
	if call.Err != nil {
		return call.Err
	}
	return nil
}
