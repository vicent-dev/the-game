package asset

import (
	_ "embed"
)

var (
	//go:embed spritesheet.png
	SpriteSheet_png []byte

	//go:embed collision.wav
	Collision_wav []byte

	//go:embed Monospace.ttf
	Monospace_ttf []byte
)
