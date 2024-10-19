package multiplayer

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"the-game/config"
)

var (
	conn *net.UDPConn
)

func connectUdpServer() *net.UDPConn {
	if conn != nil {
		return conn
	}

	c := config.LoadConfig()

	// resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", c.Server.Host+":"+c.Server.Port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// dial to the address with UDP
	conn, err = net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return conn
}

func sendServer(info []byte, variationThreshold func(data string)) {
	conn := connectUdpServer()

	_, err := conn.Write([]byte(string(info) + "\n"))

	if err != nil {
		fmt.Println(err)
		return
	}

	// read from the connection untill a new line is send
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// set the entity to server coordinates when variation is bigger than defined threshold
	variationThreshold(data)
}
