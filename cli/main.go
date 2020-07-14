package main

import (
	"fmt"

	dbus "github.com/godbus/dbus/v5"
	"github.com/shibumi/go-iwd"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	i := iwd.New(conn)
	fmt.Println(i.Networks)
	// all := conn.Object("net.connman.iwd", "/")
	// var managed map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	// err = all.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&managed)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(managed)
	iwdDevice := conn.Object("net.connman.iwd", "/net/connman/iwd/0/4")
	var networks []interface{}
	iwdDevice.Call("net.connman.iwd.Station.Scan", 0)
	err = iwdDevice.Call("net.connman.iwd.Station.GetOrderedNetworks", 0).Store(&networks)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(networks)
	// for _, network := range networks {
	// 	n := conn.Object("net.connman.iwd", network.ObjectPath)
	// 	fmt.Println(n.GetProperty("net.connman.iwd.Network.Name"))
	// }
	// //var list []string
	// //call := iwdDevice.Call("net.connman.iwd.Station.GetOrderedNetworks", 0, "")
	// //fmt.Println(call)
}
