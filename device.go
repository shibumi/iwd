package iwd

import "github.com/godbus/dbus/v5"

const (
	objectDevice      = "net.connman.iwd.Device"
	objectDeviceStart = "net.connman.iwd.Device.Start"
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
	obj := conn.Object(objectDevice, d.Path)
	call := obj.Call(callApActivate, 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}
