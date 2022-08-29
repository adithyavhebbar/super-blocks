package structs

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	Pos      rl.Vector2
	Speed    rl.Vector2
	Radius   float32
	IsActive bool
}
