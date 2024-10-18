package main

import (
	"encoding/json"
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

	// read - write
	for {
		buf := make([]byte, 512)
		rl, addr, err := conn.ReadFromUDP(buf)

		if err != nil {
			alog.Error(err.Error())
			return
		}

		match := newMatch()

		err = json.Unmarshal(buf[:rl], match)

		if err != nil {
			alog.Error("error unmarshal match ", err.Error())
		}

		if !isAddressAdded(addrs, addr) {
			addrs = append(addrs, addr)
			match.joinMatch()

			b, err := json.Marshal(match)

			if err != nil {
				alog.Error(err.Error())
			}

			alog.Info("sent ", string(b))
			go conn.WriteToUDP([]byte(string(b)+"\n"), addr)
			alog.Info("Number addrss ", len(addrs))
		} else {
			for _, a := range addrs {
				go conn.WriteToUDP([]byte(string(buf)+"\n"), a)
			}
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
