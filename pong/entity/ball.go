package entity

import (
	"math"

	"github.com/en-vee/alog"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	*Entity
}

const (
	initialAcc = 1.000000005
	initialv   = 2
)

func NewBall(sprite *ebiten.Image, x, y float64) *Ball {
	size := 7.
	scale := 4.

	e := &Entity{
		id:     "",
		sprite: sprite,
		x:      x - size*scale,
		y:      y - size*scale,
		vx:     initialv,
		vy:     initialv,
		maxvx:  8,
		maxvy:  8,
		acc:    initialAcc,
		width:  size,
		height: size,
		scale:  scale,
	}

	return &Ball{
		Entity: e,
	}
}

func (b *Ball) Update(maxX, maxY float64, gameEntities []*Entity) {
	if (b.x+(b.width*b.scale)) >= maxX || b.x < 0 {
		//@todo save score in game object
		b.x = maxX/2 - b.width*b.scale
		b.y = maxY/2 - b.height*b.scale

		b.acc = initialAcc
		if b.vx > 0 {
			b.vx = -initialv
			alog.Info("Point player")
		} else {
			b.vx = initialv
			alog.Info("Point opponent")
		}

		b.vy = initialv
	}

	if (b.y+(b.height*b.scale)) >= maxY || b.y < 0 {
		b.vy = -b.vy
	}

	if math.Abs(b.vx) <= b.maxvx {
		b.vx *= b.acc
	}

	if math.Abs(b.vy) <= b.maxvy {
		b.vy *= b.acc
	}

	if b.Collides(gameEntities) {
		b.vx = -b.vx
	}

	b.x += b.vx
	b.y += b.vy

	b.updateHitBox()
}
