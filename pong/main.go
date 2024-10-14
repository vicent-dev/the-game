package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGame()

	ebiten.SetWindowSize(g.Layout(600, 400))
	ebiten.SetWindowTitle("PONG")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
