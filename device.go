package iwd

import "github.com/godbus/dbus/v5"

const (
	objectDevice = "net.connman.iwd.Device"
)

// Device refers to the iwd network device like "wlan0" for example: /net/connman/iwd/0/4
type Device struct {
	Path    dbus.ObjectPath
	Adapter dbus.ObjectPath
	Address string
	Mode    string
	Name    string
	Powered bool
}

func (d Device) ActivateAp(conn *dbus.Conn) error {
	obj := conn.Object(objectAp, d.Path)
	call := obj.Call(callStationScan, 0, "5g", "afoe11afoe11")
	if call.Err != nil {
		return call.Err
	}
	return nil
}
