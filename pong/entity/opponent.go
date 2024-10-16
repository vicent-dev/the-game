package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Opponent struct {
	id string
	*Entity
}

func NewOpponent(id string, sprite *ebiten.Image, x, y float64) *Opponent {
	height := 32.
	width := 7.

	scale := 4.

	e := &Entity{
		id:     id,
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
