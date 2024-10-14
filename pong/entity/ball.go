package entity

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	*Entity
}

func NewBall(sprite *ebiten.Image, x, y float64) *Ball {
	size := 7.
	scale := 4.

	e := &Entity{
		sprite: sprite,
		x:      x - size*scale,
		y:      y - size*scale,
		vx:     2,
		vy:     2,
		maxvx:  8,
		maxvy:  8,
		acc:    1.0005,
		width:  size,
		height: size,
		scale:  scale,
	}

	return &Ball{
		Entity: e,
	}
}

func (b *Ball) Update(maxX, maxY float64, hitBoxes []image.Rectangle) {
	if (b.x+(b.width*b.scale)) >= maxX || b.x < 0 {
		b.vx = -b.vx
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

	for _, hb := range hitBoxes {
		if hb.Overlaps(b.HitBox) {
			b.vx = -b.vx
		}
	}

	b.x += b.vx
	b.y += b.vy

	b.updateHitBox()
}

func (b *Ball) PositionMessage() []byte {
	return []byte("ball:" + fmt.Sprintf("%f", b.x) + "," + fmt.Sprintf("%f", b.y))
}

func (b *Ball) ProcessMultiplayerResponse(data string) {
	b.Entity.processMultiplayerResponse("ball", data)
}
