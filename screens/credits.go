package screens

import (
	screenType "github.com/adithyavhebbar/blocks/gamescreen"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreditsUpdate() {
	screenType.CurrentScreenType = screenType.ScreenType.Credits
	CreditsDraw()
	CreaditsInput()
}

func CreditsDraw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawText("Adithya V Hebbar", int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2-30), 50, rl.Black)
	rl.DrawText("Software Engineer", int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2+30), 50, rl.Black)

	rl.DrawText("Q to go to menu menu", 30, 30, 30, rl.Black)
	rl.EndDrawing()
}

func CreaditsInput() {
	if rl.IsKeyPressed(rl.KeyQ) {
		screenType.CurrentScreenType = screenType.ScreenType.Menu
	}
}
