package gosnmp

import "time"

type SNMPDevice struct {
	GoSNMP
	Config *Config
}
type Config struct {
	// Port is a udp port
	Port uint16

	// Community is an SNMP Community string
	Community string

	// Version is an SNMP Version
	Version SnmpVersion

	// Timeout is the timeout for the SNMP Query
	Timeout time.Duration

	// Set the number of retries to attempt within timeout.
	Retries int
}

var DefaultConfig = &Config{
	Port:      161,
	Community: "public",
	Version:   Version2c,
	Timeout:   time.Duration(2) * time.Second,
	Retries:   3,
}

func NewSNMPDevice(cfg *Config) *SNMPDevice {

	d := &SNMPDevice{GoSNMP: GoSNMP{
		Port:      cfg.Port,
		Community: cfg.Community,
		Version:   cfg.Version,
		Timeout:   cfg.Timeout,
		Retries:   cfg.Retries,
	}, Config: cfg}
	return d
}
