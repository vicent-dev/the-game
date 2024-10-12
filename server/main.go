package main

import (
	"net"
	"os"

	"the-game-server/config"

	"github.com/en-vee/alog"
)

func main() {

	c := config.LoadConfig()

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

	// read - write
	for {
		// fixed size slice to prevent infinite reading of buffer. buf[0:] is a "fake cast" to []byte
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			alog.Error(err.Error())
			return
		}

		conn.WriteToUDP([]byte(string(buf[0:])+"\n"), addr)
	}
}
