package udp

import (
	"encoding/json"
	"net"
	"the-game-server/config"
	"the-game-server/match"

	"github.com/en-vee/alog"
)

func Serve(c *config.Config) error {
	updAddrStr := c.Server.Host + ":" + c.Server.Port

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", updAddrStr)

	if err != nil {
		return err
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		return err
	}

	alog.Info("Waiting for UDP requests [%s]", updAddrStr)

	var addrs []*net.UDPAddr
	m := match.NewMatch()

	// read - write
	for {
		buf := make([]byte, 512)
		_, addr, err := conn.ReadFromUDP(buf)

		if err != nil {
			return err
		}

		if !isAddressAdded(addrs, addr) {
			addrs = append(addrs, addr)
			m.JoinMatch()

			b, err := json.Marshal(m)

			if err != nil {
				alog.Error(err.Error())
			}

			alog.Info("send ", string(b))

			// when opponent and player are ready we cast to all the match data
			if m.Opponent.Id != "" && m.Player.Id != "" {
				broadcastAll(conn, addrs, b)
			} else {
				go conn.WriteToUDP([]byte(string(b)+"\n"), addr)
			}
		} else {
			broadcastAll(conn, addrs, buf)
		}
	}
}

func broadcastAll(conn *net.UDPConn, addrs []*net.UDPAddr, b []byte) {
	for _, a := range addrs {
		go conn.WriteToUDP([]byte(string(b)+"\n"), a)
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
