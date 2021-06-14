package iwd

import (
	"fmt"

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
	Ap            []Ap
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
		make([]Ap, 0),
	}
	for k, v := range objects {
		for resource, obj := range v {
			switch resource {
			case objectAdapter:
				i.Adapters = append(i.Adapters, Adapter{
					Path:  k,
					Model: asString(obj["Model"]), Name: asString(obj["Name"]),
					Powered: obj["Powered"].Value().(bool), SupportedModes: obj["SupportedModes"].Value().([]string),
					Vendor: asString(obj["Vendor"]),
				})
			case objectKnownNetwork:
				i.KnownNetworks = append(i.KnownNetworks, KnownNetwork{
					Path:        k,
					AutoConnect: obj["AutoConnect"].Value().(bool), Hidden: obj["Hidden"].Value().(bool),
					LastConnectedTime: asString(obj["LastConnectedTime"]), Name: asString(obj["Name"]),
					Type: asString(obj["Type"]),
				})
			case objectNetwork:
				i.Networks = append(i.Networks, Network{
					Path:      k,
					Connected: obj["Connected"].Value().(bool), Device: asPath(obj["Device"]),
					Name: asString(obj["Name"]), Type: asString(obj["Type"]),
				})
			case objectStation:
				i.Stations = append(i.Stations, Station{
					Path:             k,
					ConnectedNetwork: asPath(obj["ConnectedNetwork"]), Scanning: obj["Scanning"].Value().(bool),
					State: asString(obj["State"]),
				})
			case objectDevice:
				i.Devices = append(i.Devices, Device{
					Path:    k,
					Adapter: asPath(obj["Adapter"]), Address: asString(obj["Address"]),
					Mode: asString(obj["Mode"]), Name: asString(obj["Name"]), Powered: obj["Powered"].Value().(bool),
				})
			case objectAp:
				i.Ap = append(i.Ap, Ap{
					Path:    k,
					Started: obj["Started"].Value().(bool),
				})
				fmt.Printf("Access point : %v \n", obj)
				panic("Bonjour")
			default:
				fmt.Printf("Ressource = %v\n", resource)
			}
		}
	}
	return i
}

func asString(value dbus.Variant) string {
	if value.Value() != nil {
		return value.Value().(string)
	}
	return ""
}

func asPath(value dbus.Variant) (path dbus.ObjectPath) {
	if value.Value() != nil {
		path = value.Value().(dbus.ObjectPath)
	}
	return
}
