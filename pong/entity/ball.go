package entity

import (
	"fmt"
	"math"
)

type Ball struct {
	*Entity
}

func (b *Ball) Move(maxX, maxY float64) {
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

	b.x += b.vx
	b.y += b.vy
}

func (b *Ball) PositionMessage() []byte {
	return []byte("ball:" + fmt.Sprintf("%f", b.x) + "," + fmt.Sprintf("%f", b.y))
}

func (b *Ball) ProcessMultiplayerResponse(data string) {
	b.Entity.processMultiplayerResponse("ball", data)
}
