package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"the-game/asset"
	"the-game/entity"
	"the-game/multiplayer"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	playerRect image.Rectangle = image.Rect(0, 0, 7, 32)
	ballRect   image.Rectangle = image.Rect(14, 0, 21, 8)
	sWidth                     = 600
	sHeight                    = 400
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

	g.ball = &entity.Ball{}
	g.ball.Entity = entity.NewEntity(g.spriteSheet.SubImage(ballRect).(*ebiten.Image))

	return g
}

func (g *Game) Update() error {
	g.ball.Move(float64(sWidth), float64(sHeight))

	go multiplayer.SendServer(g.ball.PositionMessage(), g.ball.ProcessMultiplayerResponse)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}
