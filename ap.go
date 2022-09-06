package iwd

import "github.com/godbus/dbus/v5"

const (
	objectAp       = "net.connman.iwd.AccessPoint"
	callApActivate = "net.connman.iwd.AccessPoint.StartProfile"
	callStop       = "net.connman.iwd.AccessPoint.Stop"
)

// Station refers to net.connman.iwd.Station
type Ap struct {
	Path    dbus.ObjectPath
	Started bool
}

func (a Ap) SetFile(conn *dbus.Conn, file string) error {
	obj := conn.Object(objectIwd, a.Path)
	call := obj.Call(callApActivate, 0, file)
	return call.Err
}

func (a Ap) Stop(conn *dbus.Conn) error {
	obj := conn.Object(objectIwd, a.Path)
	call := obj.Call(callStop, 0)
	return call.Err
}
