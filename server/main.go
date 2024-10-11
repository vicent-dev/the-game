package main

import (
	"fmt"
	"net"
	"os"

	"github.com/en-vee/alog"
)

func main() {

	c := LoadConfig()

	updAddrStr := c.Server.Host + ":" + c.Server.Port

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", updAddrStr)

	if err != nil {
		alog.Error(err.Error())
		os.Exit(1)
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		alog.Error(err.Error())
		os.Exit(1)
	}

	alog.Info("Waiting for UDP requests [%s]", updAddrStr)

	// Read from UDP listener in endless loop
	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			alog.Error(err.Error())
			return
		}

		fmt.Print("> ", string(buf[0:]))

		// Write back the message over UPD
		conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
	}
}
