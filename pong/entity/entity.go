package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	sprite *ebiten.Image
	x      float64
	y      float64
	vx     float64
	vy     float64
	maxvx  float64
	maxvy  float64
	acc    float64
	width  float64
	height float64
	scale  float64
}

func NewEntity(sprite *ebiten.Image) *Entity {
	return &Entity{
		sprite: sprite,
		x:      0,
		y:      0,
		vx:     2,
		vy:     2,
		maxvx:  8,
		maxvy:  8,
		acc:    1.0005,
		width:  7,
		height: 7,
		scale:  4,
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Scale(e.scale, e.scale)
	opts.GeoM.Translate(e.x, e.y)

	screen.DrawImage(e.sprite, opts)
}

func (e *Entity) GetBounds() (float64, float64) {
	return e.x, e.y
}
