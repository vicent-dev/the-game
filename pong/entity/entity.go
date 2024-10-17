package entity

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const variationThreshold = 4.0001

type Entity struct {
	id     string
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
	hitBox image.Rectangle
}

func (e *Entity) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Scale(e.scale, e.scale)
	opts.GeoM.Translate(e.x, e.y)

	screen.DrawImage(e.sprite, opts)
}

func (e *Entity) updateHitBox() {
	e.hitBox = image.Rect(
		int(e.x),
		int(e.y),
		int(e.x+e.width*e.scale),
		int(e.y+e.height*e.scale),
	)
}

func (e *Entity) possitionMessage(key string) []byte {
	return []byte(key + "-" + e.id + ":" + fmt.Sprintf("%f", e.x) + "," + fmt.Sprintf("%f", e.y))
}

func (e *Entity) ProcessMultiplayerCoordinates(x, y float64) {
	if math.Abs(x-e.x) >= variationThreshold || math.Abs(y-e.y) >= variationThreshold {
		e.x = x
		e.y = y
	}
}

func (e *Entity) Coordinates() (float64, float64) {
	return e.x, e.y
}

func (e *Entity) Collides(entities []*Entity) bool {
	for _, entity := range entities {
		if entity.hitBox.Overlaps(e.hitBox) {
			return true
		}
	}

	return false
}

func (e *Entity) PrintHitBox(screen *ebiten.Image) {
	vector.StrokeRect(screen,
		float32(e.hitBox.Min.X),
		float32(e.hitBox.Min.Y),
		float32(e.hitBox.Dx()),
		float32(e.hitBox.Dy()),
		3.0,
		color.RGBA{255, 0, 0, 255},
		true,
	)
}
