package screens

import (
	"fmt"

	screenType "github.com/adithyavhebbar/blocks/gamescreen"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var MousePosition rl.Vector2

var PlayRect rl.Rectangle
var CreditsRect rl.Rectangle
var ControllsRect rl.Rectangle
var ExitRect rl.Rectangle

func MenuInit() {
	PlayRect.Height = float32(rl.GetScreenHeight()) / 10
	PlayRect.Width = float32(rl.GetScreenWidth()) / 10
	PlayRect.X = float32(rl.GetScreenWidth())/9 - PlayRect.Width
	PlayRect.Y = float32(rl.GetScreenHeight())/9 - PlayRect.Height

	CreditsRect.Height = float32(rl.GetScreenHeight()) / 10
	CreditsRect.Width = float32(rl.GetScreenWidth()) / 10
	CreditsRect.X = float32(rl.GetScreenWidth())/9 - PlayRect.Width
	CreditsRect.Y = float32(rl.GetScreenHeight())/9 + PlayRect.Height

	ControllsRect.Height = float32(rl.GetScreenHeight()) / 10
	ControllsRect.Width = float32(rl.GetScreenWidth()) / 10
	ControllsRect.X = float32(rl.GetScreenWidth())/9 - PlayRect.Width
	ControllsRect.Y = float32(rl.GetScreenHeight())/9 + (PlayRect.Height * 3)

	ExitRect.Height = float32(rl.GetScreenHeight()) / 10
	ExitRect.Width = float32(rl.GetScreenWidth()) / 10
	ExitRect.X = float32(rl.GetScreenWidth())/9 - PlayRect.Width
	ExitRect.Y = float32(rl.GetScreenHeight())/9 + (PlayRect.Height * 5)

}

func MenuUpdate() {
	MenuDraw()

	MousePosition = rl.GetMousePosition()

	if rl.CheckCollisionPointRec(MousePosition, PlayRect) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			fmt.Println("Play Pressed")
			screenType.CurrentScreenType = screenType.ScreenType.Game
		}
	}

	if rl.CheckCollisionPointRec(MousePosition, CreditsRect) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			fmt.Println("Credits Pressed")
			screenType.CurrentScreenType = screenType.ScreenType.Credits
		}
	}

	if rl.CheckCollisionPointRec(MousePosition, ControllsRect) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			fmt.Println("Controls Pressed")
			screenType.CurrentScreenType = screenType.ScreenType.Controls
		}
	}

	if rl.CheckCollisionPointRec(MousePosition, ExitRect) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			fmt.Println("Exit Pressed")
			screenType.CurrentScreenType = screenType.ScreenType.End
		}
	}
}

func MenuDraw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.DrawRectangle(int32(PlayRect.X), int32(PlayRect.Y), int32(PlayRect.Width), int32(PlayRect.Height), rl.Black)

	rl.DrawRectangle(int32(CreditsRect.X), int32(CreditsRect.Y), int32(CreditsRect.Width), int32(CreditsRect.Height), rl.Black)

	rl.DrawRectangle(int32(ControllsRect.X), int32(ControllsRect.Y), int32(ControllsRect.Width), int32(ControllsRect.Height), rl.Black)

	rl.DrawRectangle(int32(ExitRect.X), int32(ExitRect.Y), int32(ExitRect.Width), int32(ExitRect.Height), rl.Black)

	// Draw Text on Rectangle
	xCenter := (PlayRect.ToInt32().Width / 2) - (rl.MeasureText("Play", 20) / 2)
	yCenter := int32(PlayRect.Height / 2)
	rl.DrawText("Play", PlayRect.ToInt32().X+xCenter, int32(PlayRect.Y+float32(yCenter)), 20, rl.RayWhite)

	xCenter = (CreditsRect.ToInt32().Width / 2) - (rl.MeasureText("Credits", 20) / 2)
	yCenter = int32(CreditsRect.Height / 2)
	rl.DrawText("Credits", CreditsRect.ToInt32().X+xCenter, int32(CreditsRect.Y+float32(yCenter)), 20, rl.RayWhite)

	xCenter = (ControllsRect.ToInt32().Width / 2) - (rl.MeasureText("Controlls", 20) / 2)
	yCenter = int32(ControllsRect.Height / 2)
	rl.DrawText("Controlls", ControllsRect.ToInt32().X+xCenter, int32(ControllsRect.Y+float32(yCenter)), 20, rl.RayWhite)

	xCenter = (ExitRect.ToInt32().Width / 2) - (rl.MeasureText("Exit", 20) / 2)
	yCenter = int32(ExitRect.Height / 2)
	rl.DrawText("Exit", ExitRect.ToInt32().X+xCenter, int32(ExitRect.Y+float32(yCenter)), 20, rl.RayWhite)

	rl.DrawText("Super", int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2-50), 100, rl.Red)
	rl.DrawText("Blocks", int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2)+40, 100, rl.Red)
	rl.DrawText("Use the mouse to move and Interact", 10, int32(rl.GetScreenHeight()-40), 40, rl.Black)
	rl.EndDrawing()
}
