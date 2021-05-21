package iwd

import (
	"github.com/godbus/dbus/v5"
)

const (
	objectIwd              = "net.connman.iwd"
	objectIwdPath          = "/net/connman/iwd"
	iwdAgentManager        = "net.connman.iwd.AgentManager"
	iwdAdapter             = "net.connman.iwd.Adapter"
	iwdDevice              = "net.connman.iwd.Device"
	iwdSimpleConfiguration = "net.connman.iwd.SimpleConfiguation"
	iwdNetwork             = "net.connman.iwd.Network"
)

// Iwd is a struct over all major iwd components
type Iwd struct {
	Agents        []Agent
	Adapters      []Adapter
	KnownNetworks []KnownNetwork
	Networks      []Network
	Stations      []Station
	Devices       []Device
}

// New parses the net.connman.iwd object index and initializes an iwd object
func New(conn *dbus.Conn) Iwd {
	var objects map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	objectManager := conn.Object(objectIwd, "/")
	objectManager.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&objects)
	i := Iwd{
		make([]Agent, 0),
		make([]Adapter, 0),
		make([]KnownNetwork, 0),
		make([]Network, 0),
		make([]Station, 0),
		make([]Device, 0),
	}
	for k, v := range objects {
		for resource, obj := range v {
			switch resource {
			case objectAdapter:
				model := ""
				name := ""
				if obj["Model"].Value() == nil {
					model = ""
				} else {
					model = obj["Model"].Value().(string)
				}
				if obj["Name"].Value() == nil {
					name = ""
				} else {
					name = obj["Name"].Value().(string)
				}
				i.Adapters = append(i.Adapters, Adapter{
					Path:  k,
					Model: model, Name: name,
					Powered: obj["Powered"].Value().(bool), SupportedModes: obj["SupportedModes"].Value().([]string),
					Vendor: obj["Vendor"].Value().(string),
				})
			case objectKnownNetwork:
				i.KnownNetworks = append(i.KnownNetworks, KnownNetwork{
					Path:        k,
					AutoConnect: obj["AutoConnect"].Value().(bool), Hidden: obj["Hidden"].Value().(bool),
					LastConnectedTime: obj["LastConnectedTime"].Value().(string), Name: obj["Name"].Value().(string),
					Type: obj["Type"].Value().(string),
				})
			case objectNetwork:
				i.Networks = append(i.Networks, Network{
					Path:      k,
					Connected: obj["Connected"].Value().(bool), Device: obj["Device"].Value().(dbus.ObjectPath),
					Name: obj["Name"].Value().(string), Type: obj["Type"].Value().(string),
				})
			case objectStation:
				i.Stations = append(i.Stations, Station{
					Path:             k,
					ConnectedNetwork: obj["ConnectedNetwork"].Value().(dbus.ObjectPath), Scanning: obj["Scanning"].Value().(bool),
					State: obj["State"].Value().(string),
				})
			case objectDevice:
				i.Devices = append(i.Devices, Device{
					Path:    k,
					Adapter: obj["Adapter"].Value().(dbus.ObjectPath), Address: obj["Address"].Value().(string),
					Mode: obj["Mode"].Value().(string), Name: obj["Name"].Value().(string), Powered: obj["Powered"].Value().(bool),
				})
			}
		}
	}
	return i
}
