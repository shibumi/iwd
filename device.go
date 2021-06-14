package iwd

import "github.com/godbus/dbus/v5"

const (
	objectDevice = "net.connman.iwd.Device"
	object       = "net.connman.iwd"
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

func (d Device) SetApOn(conn *dbus.Conn) error {
	obj := conn.Object(object, d.Path)
	err := obj.SetProperty(objectDevice+".Mode", dbus.MakeVariant("ap"))
	if err != nil {
		return err
	}
	err = obj.SetProperty(objectDevice+".Powered", dbus.MakeVariant("on"))
	if err != nil {
		return err
	}
	return nil
}
