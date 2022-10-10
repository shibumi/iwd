package iwd

import (
	"github.com/godbus/dbus/v5"
)

const (
	callStationGetDiagnostics = "net.connman.iwd.StationDiagnostic.GetDiagnostics"
	callStationScan           = "net.connman.iwd.Station.Scan"
	objectStation             = "net.connman.iwd.Station"
)

// Station refers to net.connman.iwd.Station
type Station struct {
	Path             dbus.ObjectPath
	ConnectedNetwork dbus.ObjectPath
	Scanning         bool
	State            string
}

// StationDiagnostic refers to results from net.connman.iwd.StationDiagnostic.GetDiagnostics
type StationDiagnostic struct {
	AverageRSSI  int16
	ConnectedBss string
	Frequency    uint32
	RSSI         int16
	RxBitrate    uint32
	RxMCS        uint8
	RxMode       string
	Security     string
	TxBitrate    uint32
	TxMCS        uint8
	TxMode       string
}

// Scan scans for wireless networks
func (s *Station) Scan(conn *dbus.Conn) error {
	obj := conn.Object(objectStation, "")
	call := obj.Call(callStationScan, 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}

// Gather StationDiagnostics for this Station
func (s *Station) GetDiagnostics(conn *dbus.Conn) (*StationDiagnostic, error) {
	var diagnostics map[string]dbus.Variant
	err := conn.Object(objectIwd, s.Path).Call(callStationGetDiagnostics, 0).Store(&diagnostics)
	if err != nil {
		return nil, err
	}
	return &StationDiagnostic{
		AverageRSSI:  diagnostics["AverageRSSI"].Value().(int16),
		ConnectedBss: diagnostics["ConnectedBss"].Value().(string),
		Frequency:    diagnostics["Frequency"].Value().(uint32),
		RSSI:         diagnostics["RSSI"].Value().(int16),
		RxBitrate:    diagnostics["RxBitrate"].Value().(uint32),
		RxMCS:        diagnostics["RxMCS"].Value().(uint8),
		RxMode:       diagnostics["RxMode"].Value().(string),
		Security:     diagnostics["Security"].Value().(string),
		TxBitrate:    diagnostics["TxBitrate"].Value().(uint32),
		TxMCS:        diagnostics["TxMCS"].Value().(uint8),
		TxMode:       diagnostics["TxMode"].Value().(string),
	}, nil
}
