package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"the-game/asset"
	"the-game/entity"
	"the-game/multiplayer"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	playerRect   image.Rectangle = image.Rect(0, 0, 7, 32)
	ballRect     image.Rectangle = image.Rect(14, 0, 21, 8)
	opponentRect image.Rectangle = image.Rect(7, 0, 14, 32)

	sWidth  = 600
	sHeight = 400
)

type Game struct {
	spriteSheet *ebiten.Image
	player      *entity.Player
	opponent    *entity.Opponent
	ball        *entity.Ball
}

func NewGame() *Game {
	g := &Game{}

	img, _, err := image.Decode(bytes.NewReader(asset.SpriteSheet_png))

	if err != nil {
		log.Fatal(err.Error())
	}

	g.spriteSheet = ebiten.NewImageFromImage(img)

	g.ball = entity.NewBall(g.spriteSheet.SubImage(ballRect).(*ebiten.Image), float64(sWidth)/2, float64(sHeight)/2)
	g.player = entity.NewPlayer(g.spriteSheet.SubImage(playerRect).(*ebiten.Image), float64(sHeight)/2)
	g.opponent = entity.NewOpponent(g.spriteSheet.SubImage(opponentRect).(*ebiten.Image), float64(sWidth), float64(sHeight)/2)

	return g
}

func (g *Game) Update() error {
	w, h := float64(sWidth), float64(sHeight)

	g.player.Update(w, h)

	go multiplayer.SendServer(g.player.PositionMessage(), g.player.ProcessMultiplayerResponse)

	g.opponent.Update(w, h)

	go multiplayer.SendServer(g.opponent.PositionMessage(), g.opponent.ProcessMultiplayerResponse)

	g.ball.Update(w, h, []image.Rectangle{
		g.player.HitBox,
		g.opponent.HitBox,
	})

	go multiplayer.SendServer(g.ball.PositionMessage(), g.ball.ProcessMultiplayerResponse)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.player.Draw(screen)
	g.opponent.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func (g *Game) printHitboxes(screen *ebiten.Image) {
	vector.StrokeRect(screen,
		float32(g.ball.HitBox.Min.X),
		float32(g.ball.HitBox.Min.Y),
		float32(g.ball.HitBox.Dx()),
		float32(g.ball.HitBox.Dy()),
		3.0,
		color.RGBA{255, 0, 0, 255},
		true,
	)

	vector.StrokeRect(screen,
		float32(g.player.HitBox.Min.X),
		float32(g.player.HitBox.Min.Y),
		float32(g.player.HitBox.Dx()),
		float32(g.player.HitBox.Dy()),
		3.0,
		color.RGBA{0, 255, 0, 255},
		true,
	)

	vector.StrokeRect(screen,
		float32(g.opponent.HitBox.Min.X),
		float32(g.opponent.HitBox.Min.Y),
		float32(g.opponent.HitBox.Dx()),
		float32(g.opponent.HitBox.Dy()),
		3.0,
		color.RGBA{0, 0, 255, 255},
		true,
	)
}
