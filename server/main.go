package main

import (
	"os"
	"the-game-server/config"
	"the-game-server/udp"

	"github.com/en-vee/alog"
)

func main() {

	c := config.LoadConfig()

	if err := udp.Serve(c); err != nil {
		alog.Error(err.Error())
		os.Exit(1)
	}
}
