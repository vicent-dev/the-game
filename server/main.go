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

	var addrs []*net.UDPAddr

	// @todo change to have multiple matches for a single server instance:w
	match := &match{}

	// read - write
	for {
		// fixed size slice to prevent infinite reading of buffer. buf[0:] is a "fake cast" to []byte
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])

		if err != nil {
			alog.Error(err.Error())
			return
		}

		if !isAddressAdded(addrs, addr) {
			addrs = append(addrs, addr)
			match.joinMatch()
		}

		alog.Info("Number addrss ", len(addrs))

		for _, a := range addrs {
			go conn.WriteToUDP([]byte(string(buf[0:])+"\n"), a)
		}
	}
}

func isAddressAdded(addrs []*net.UDPAddr, addr *net.UDPAddr) bool {
	for _, a := range addrs {
		if a.IP.Equal(addr.IP) && a.Port == addr.Port && a.Zone == addr.Zone {
			return true
		}
	}

	return false
}
