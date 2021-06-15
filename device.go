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

func (d Device) SetAp(conn *dbus.Conn) error {
	obj := conn.Object(objectIwd, d.Path)
	err := obj.SetProperty(objectDevice+".Mode", dbus.MakeVariant("ap"))
	return err
}

func (d Device) SetOn(conn *dbus.Conn) error {
	obj := conn.Object(objectIwd, d.Path)
	err := obj.SetProperty(objectDevice+".Powered", dbus.MakeVariant(true))
	return err
}

func (d Device) SetOff(conn *dbus.Conn) error {
	obj := conn.Object(objectIwd, d.Path)
	err := obj.SetProperty(objectDevice+".Powered", dbus.MakeVariant(false))
	return err
}
