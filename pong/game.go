package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"strconv"
	"the-game/asset"
	"the-game/entity"
	"the-game/multiplayer"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	playerRect image.Rectangle = image.Rect(0, 0, 7, 32)
	ballRect   image.Rectangle = image.Rect(14, 0, 21, 8)
	sHeight                    = 320
	sWidth                     = 240
)

type Game struct {
	spriteSheet *ebiten.Image
	player      *entity.Player
	ball        *entity.Ball
}

func NewGame() *Game {
	g := &Game{}

	img, _, err := image.Decode(bytes.NewReader(asset.SpriteSheet_png))

	if err != nil {
		log.Fatal(err.Error())
	}

	g.spriteSheet = ebiten.NewImageFromImage(img)

	g.player = &entity.Player{}
	g.player.Sprite = g.spriteSheet.SubImage(playerRect).(*ebiten.Image)

	g.ball = &entity.Ball{}
	g.ball.Sprite = g.spriteSheet.SubImage(ballRect).(*ebiten.Image)

	return g
}

func (g *Game) Update() error {
	cx, cy := ebiten.CursorPosition()
	multiplayer.SendServer([]byte(strconv.Itoa(cx) + "," + strconv.Itoa(cy)))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	screen.DrawImage(g.player.Sprite, opts)
	screen.DrawImage(g.ball.Sprite, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}
