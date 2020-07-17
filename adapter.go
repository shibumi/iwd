package iwd

import "github.com/godbus/dbus/v5"

const (
	objectAdapter = "net.connman.iwd.Adapter"
)

// Adapter refers to the iwd network adapter like "phy/0" for example: /net/connman/iwd/0
type Adapter struct {
	Path           dbus.ObjectPath
	Model          string
	Name           string
	Powered        bool
	SupportedModes []string
	Vendor         string
}
