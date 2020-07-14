package iwd

import "github.com/godbus/dbus/v5"

const (
	objectDevice = "net.connman.iwd.Device"
)

// Device refers to the iwd network device like "wlan0" for example: /net/connman/iwd/0/4
type Device struct {
	Adapter dbus.ObjectPath
	Address string
	Mode    string
	Name    string
	Powered bool
}
