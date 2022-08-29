package gamemanager

import (
	screenType "github.com/adithyavhebbar/blocks/gamescreen"
	screens "github.com/adithyavhebbar/blocks/screens"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 2500
	screenHeight = 1500
)

var Screen int
var playing bool

func Run() {

	screenType.SetScreenType(screenType.ScreenType.Menu)
	Screen = screenType.GetScreenType()
	playing = true
	rl.InitWindow(screenWidth, screenHeight, "Block Breaker")

	Init()

	for !rl.WindowShouldClose() && playing {
		Change()
	}

	DeInit()

	rl.CloseWindow()
}

func Init() {
	screens.MenuInit()
	screens.GameInit()
	rl.InitAudioDevice()
}

func DeInit() {
	rl.CloseAudioDevice()
}

func Change() {
	switch screenType.CurrentScreenType {
	case screenType.ScreenType.Menu:
		screens.MenuUpdate()

	case screenType.ScreenType.Game:
		screens.GamePlayUpdate()

	case screenType.ScreenType.Controls:
		break

	case screenType.ScreenType.Credits:
		break

	case screenType.ScreenType.End:
		playing = false
	}
}
