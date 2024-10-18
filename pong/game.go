package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"the-game/asset"
	"the-game/entity"
	"the-game/multiplayer"
	"time"

	"github.com/en-vee/alog"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	playerRect   image.Rectangle = image.Rect(0, 0, 7, 32)
	ballRect     image.Rectangle = image.Rect(14, 0, 21, 8)
	opponentRect image.Rectangle = image.Rect(7, 0, 14, 32)

	sWidth  = 600
	sHeight = 400
)

type Game struct {
	spriteSheet  *ebiten.Image
	match        *multiplayer.Match
	player       *entity.Player
	opponent     *entity.Opponent
	ball         *entity.Ball
	synchronized bool
	font         *text.GoTextFaceSource
}

func NewGame() *Game {
	g := &Game{
		synchronized: false,
	}

	g.match = multiplayer.NewMatch() // @todo make connection process

	s, err := text.NewGoTextFaceSource(bytes.NewReader(asset.Monospace_ttf))

	if err != nil {
		log.Fatal(err)
	}

	g.font = s

	g.loadEntities()

	return g
}

func (g *Game) loadEntities() {
	img, _, err := image.Decode(bytes.NewReader(asset.SpriteSheet_png))

	if err != nil {
		log.Fatal(err.Error())
	}

	g.spriteSheet = ebiten.NewImageFromImage(img)

	g.ball = entity.NewBall(g.spriteSheet.SubImage(ballRect).(*ebiten.Image), float64(sWidth)/2, float64(sHeight)/2)
	g.player = entity.NewPlayer("", g.spriteSheet.SubImage(playerRect).(*ebiten.Image), float64(sHeight)/2)
	g.opponent = entity.NewOpponent("", g.spriteSheet.SubImage(opponentRect).(*ebiten.Image), float64(sWidth), float64(sHeight)/2)
}

func (g *Game) Update() error {
	if !g.match.Ready() {
		g.match.JoinMatch()
		return nil
	}

	w, h := float64(sWidth), float64(sHeight)

	g.player.Update(w, h)
	g.opponent.Update(w, h)
	g.ball.Update(w, h, []*entity.Entity{
		g.player.Entity,
		g.opponent.Entity,
	})

	if false == g.synchronized && time.Now().Local().Second()%2 == 0 {
		g.synchronized = true
		go g.sync()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if !g.match.Ready() {
		op := &text.DrawOptions{}
		op.GeoM.Translate(10, 20)
		op.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, "Looking for a fair opponent...", &text.GoTextFace{
			Source: g.font,
			Size:   32,
		}, op)

		return
	}

	g.ball.Draw(screen)
	g.player.Draw(screen)
	g.opponent.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func (g *Game) sync() {
	// first sync our match entity
	g.match.Player.X, g.match.Player.Y = g.player.Coordinates()
	g.match.Opponent.X, g.match.Opponent.Y = g.opponent.Coordinates()
	g.match.Ball.X, g.match.Ball.Y = g.ball.Coordinates()
	g.match.UpdatedAt = time.Now()

	g.match.Sync(func(data string) {
		g.synchronized = false
		//save into match
		err := json.Unmarshal([]byte(data), g.match)

		if err != nil {
			alog.Error(err.Error())
			return
		}

		// sync game from match
		delay := time.Now().Sub(g.match.UpdatedAt).Seconds()

		if delay > 1 {
			alog.Error("server update not applied. Delay bigger than 1 second ", delay)
			return
		}

		g.ball.ProcessMultiplayerCoordinates(g.match.Ball.X, g.match.Ball.Y)
		g.player.ProcessMultiplayerCoordinates(g.match.Player.X, g.match.Player.Y)
		g.opponent.ProcessMultiplayerCoordinates(g.match.Opponent.X, g.match.Opponent.Y)
	})
}

func (g *Game) printHitboxes(screen *ebiten.Image) {
	g.ball.PrintHitBox(screen)
	g.player.PrintHitBox(screen)
	g.opponent.PrintHitBox(screen)
}
