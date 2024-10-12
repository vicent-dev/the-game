package main

import (
	"bytes"
	"image"
	"strconv"
	"the-game/asset"
	"the-game/multiplayer"

	"github.com/en-vee/alog"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	spriteSheet *ebiten.Image
}

func NewGame() *Game {
	g := &Game{}

	img, _, err := image.Decode(bytes.NewReader(asset.Spritesheet_png))

	if err != nil {
		alog.Error(err.Error())
		return nil
	}

	g.spriteSheet = ebiten.NewImageFromImage(img)

	return g
}

func (g *Game) Update() error {
	cx, cy := ebiten.CursorPosition()
	multiplayer.SendServer([]byte(strconv.Itoa(cx) + "," + strconv.Itoa(cy)))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
