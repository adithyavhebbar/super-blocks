package screens

import (
	screenType "github.com/adithyavhebbar/blocks/gamescreen"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ControlsUpdate() {
	screenType.CurrentScreenType = screenType.ScreenType.Controls
	ControlsDraw()
	ControlsInput()
}

func ControlsDraw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawText("Press SPACE to release the ball", int32(rl.GetScreenHeight()/2), int32(rl.GetScreenHeight()/5), 40, rl.Black)
	rl.DrawText("Press <- to move left", int32(rl.GetScreenHeight()/2), int32(rl.GetScreenHeight()/3), 40, rl.Black)

	rl.DrawText("Press -> to move right", int32(rl.GetScreenHeight()/2), int32(rl.GetScreenHeight()/2), 40, rl.Black)

	rl.DrawText("Press SPACE to release the ball", int32(rl.GetScreenHeight()/2), int32(rl.GetScreenHeight()/5), 40, rl.Black)
	rl.DrawText("Q to go to menu menu", 30, 30, 30, rl.Black)

	rl.EndDrawing()
}

func ControlsInput() {
	if rl.IsKeyPressed(rl.KeyQ) {
		screenType.CurrentScreenType = screenType.ScreenType.Menu
	}
}
