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

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", c.Server.Host+":"+c.Server.Port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err = net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return conn
}

func SendServer(info []byte) {
	// Send a message to the server
	conn := connectUdpServer()

	_, err := conn.Write([]byte(string(info) + "\n"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from the connection untill a new line is send
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data read from the connection to the terminal
	fmt.Print("> ", string(data))
}
