package main

import (
	"the-game-server/config"
	"the-game-server/udp"
)

func main() {

	c := config.LoadConfig()

	udp.Serve(c)
}
