package entity

import (
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

const variationThreshold = .00000001

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

func (e *Entity) processMultiplayerResponse(key string, data string) {
	// remove EOL
	data, _ = strings.CutSuffix(data, "\n")
	dataSplit := strings.Split(data, ":")

	if len(dataSplit) == 0 {
		return
	}

	if dataSplit[0] != key {
		return
	}

	serverCoordinates := strings.Split(dataSplit[1], ",")

	ssx := serverCoordinates[0]
	ssy := serverCoordinates[1]

	sx, err := strconv.ParseFloat(ssx, 64)

	if err != nil {
		return
	}

	sy, err := strconv.ParseFloat(ssy, 64)

	if err != nil {
		return
	}

	if math.Abs(sx-e.x) >= variationThreshold || math.Abs(sy-e.y) >= variationThreshold {
		e.x = sx
		e.y = sy
	}
}
