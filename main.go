package main

import (
	"fmt"

	dbus "github.com/godbus/dbus/v5"
	"github.com/shibumi/iwd-menu/iwd"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	iwdDevice := conn.Object("net.connman.iwd", "/net/connman/iwd/0/4")
	var networks iwd.Networks
	iwdDevice.Call("net.connman.iwd.Station.Scan", 0)
	iwdDevice.Call("net.connman.iwd.Station.GetOrderedNetworks", 0).Store(&networks)
	fmt.Println(networks)
	for _, network := range networks {
		n := conn.Object("net.connman.iwd", network.ObjectPath)
		fmt.Println(n.GetProperty("net.connman.iwd.Network.Name"))
	}
	//var list []string
	//call := iwdDevice.Call("net.connman.iwd.Station.GetOrderedNetworks", 0, "")
	//fmt.Println(call)
}
