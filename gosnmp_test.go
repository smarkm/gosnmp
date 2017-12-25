package gosnmp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/smarkm/gosnmp"
)

func TestDefaultTest(t *testing.T) {
	go synchronize()
	time.Sleep(5 * time.Second)
}

func synchronize() {
	cfg := gosnmp.DefaultConfig
	d := gosnmp.NewSNMPDevice(cfg)
	d.Target = "127.0.0.1"
	if err := d.Connect(); err != nil {
		return
	}
	defer d.Conn.Close()

	oids := []string{
		"1.3.6.1.2.1.1.1.0",
		"1.3.6.1.2.1.1.2.0",
		"1.3.6.1.2.1.1.3.0",
		"1.3.6.1.2.1.1.4.0",
		"1.3.6.1.2.1.1.5.0",
		"1.3.6.1.2.1.1.7.0",
	}
	rs, _ := d.Get(oids)
	for i, variable := range rs.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}
}
