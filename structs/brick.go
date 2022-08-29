package structs

import rl "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	Pos      rl.Vector2
	Size     rl.Vector2
	IsActive bool
	Color    int
}

func (brick *Brick) GetRect() rl.Rectangle {
	return rl.Rectangle{X: brick.Pos.X, Y: brick.Pos.Y, Width: brick.Size.X, Height: brick.Size.Y}
}
