package structs

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Speed int
	Lives int
}

func (player *Player) GetRect() rl.Rectangle {
	return rl.Rectangle{X: player.Pos.X, Y: player.Pos.Y, Width: player.Size.X, Height: player.Size.Y}
}
