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
	spriteSheet  *ebiten.Image
	match        *multiplayer.Match
	player       *entity.Player
	opponent     *entity.Opponent
	ball         *entity.Ball
	synchronized bool
}

func NewGame() *Game {
	g := &Game{
		synchronized: false,
	}

	img, _, err := image.Decode(bytes.NewReader(asset.SpriteSheet_png))

	if err != nil {
		log.Fatal(err.Error())
	}

	g.match = multiplayer.NewMatch() // @todo make connection process

	g.spriteSheet = ebiten.NewImageFromImage(img)

	g.ball = entity.NewBall(g.spriteSheet.SubImage(ballRect).(*ebiten.Image), float64(sWidth)/2, float64(sHeight)/2)
	g.player = entity.NewPlayer("", g.spriteSheet.SubImage(playerRect).(*ebiten.Image), float64(sHeight)/2)
	g.opponent = entity.NewOpponent("", g.spriteSheet.SubImage(opponentRect).(*ebiten.Image), float64(sWidth), float64(sHeight)/2)

	return g
}

func (g *Game) Update() error {
	w, h := float64(sWidth), float64(sHeight)

	g.player.Update(w, h)
	g.opponent.Update(w, h)
	g.ball.Update(w, h, []image.Rectangle{
		g.player.HitBox,
		g.opponent.HitBox,
	})

	if false == g.synchronized && time.Now().Local().Second()%2 == 0 {
		g.synchronized = true
		go g.sync()
	}

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
