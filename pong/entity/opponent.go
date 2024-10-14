package entity

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Opponent struct {
	*Entity
}

func NewOpponent(sprite *ebiten.Image, x, y float64) *Opponent {
	height := 32.
	width := 7.

	scale := 4.

	e := &Entity{
		sprite: sprite,
		x:      x - 10 - (width * scale),
		y:      y - (height*scale)/2,
		vx:     0, // not needed
		vy:     3,
		maxvx:  0, // not needed
		maxvy:  3,
		acc:    0,
		width:  width,
		height: height,
		scale:  scale,
	}

	return &Opponent{
		Entity: e,
	}
}

func (o *Opponent) Update(maxX, maxY float64) {
	o.updateHitBox()
}

func (o *Opponent) PositionMessage() []byte {
	return []byte("opponent:" + fmt.Sprintf("%f", o.x) + "," + fmt.Sprintf("%f", o.y))
}

func (o *Opponent) ProcessMultiplayerResponse(data string) {
	o.Entity.processMultiplayerResponse("opponent", data)

	// @todo override default logic and trust server data
}
