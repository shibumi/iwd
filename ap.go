package iwd

import "github.com/godbus/dbus/v5"

const (
	objectAp       = "net.connman.iwd.AccessPoint"
	callApActivate = "net.connman.iwd.AccessPoint.StartProfile"
)

// Station refers to net.connman.iwd.Station
type Ap struct {
	Path    dbus.ObjectPath
	Started bool
}

func (a Ap) Activateconn(conn *dbus.Conn) error {
	obj := conn.Object(objectAp, "")
	call := obj.Call(callApActivate, 0, "monsuperwifi")
	if call.Err != nil {
		return call.Err
	}
	return nil
}
