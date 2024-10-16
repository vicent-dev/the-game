package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*Entity
}

func NewPlayer(id string,sprite *ebiten.Image, y float64) *Player {
	height := 32.
	scale := 4.
	e := &Entity{
		id:     id,
		sprite: sprite,
		x:      10,
		y:      y - (height*scale)/2,
		vx:     0, // not needed
		vy:     3,
		maxvx:  0, // not needed
		maxvy:  3,
		acc:    0,
		width:  7,
		height: height,
		scale:  scale,
	}

	return &Player{
		Entity: e,
	}
}

func (p *Player) Update(maxX, maxY float64) {
	if p.y > 0 && ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		p.y -= p.vy
	} else if (p.y+p.height*p.scale) <= maxY && ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		p.y += p.vy
	}

	p.updateHitBox()
}