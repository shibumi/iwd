package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/shibumi/iwd"
)

// This little example shows the network name of the connected wifi network.
func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	iwdClient := iwd.New(conn)
	// lookup connected network
	var networkPath dbus.ObjectPath
	for _, station := range iwdClient.Stations {
		if station.State == "connected" {
			networkPath = station.ConnectedNetwork
			break
		}
	}
	for _, network := range iwdClient.Networks {
		if network.Path == networkPath {
			fmt.Println(network.Name)
		}
	}
}
