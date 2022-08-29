package screens

import (
	screenType "github.com/adithyavhebbar/blocks/gamescreen"
	structs "github.com/adithyavhebbar/blocks/structs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BrickRows    = 4
	BrickColumns = 10
	Offset       = 80
)

var Player structs.Player
var Bricks [BrickRows][BrickColumns]structs.Brick
var Ball structs.Ball
var Points int

var Poing rl.Sound
var Music rl.Music

var win bool = false
var pause bool = false

var Color [4]rl.Color = [4]rl.Color{rl.Black, rl.Gray, rl.DarkGreen, rl.Magenta}

func GameInit() {
	Points = 0
	Poing = rl.LoadSound("assets/Player_clition.wav")
	Music = rl.LoadMusicStream("assets/melodic-techno-03-extended-version-moogify-9867.mp3")

	Player.Size = rl.Vector2{X: float32(rl.GetScreenWidth() / 10), Y: float32(rl.GetScreenHeight() / 30)}
	Player.Pos = rl.Vector2{X: float32(rl.GetScreenWidth()/2 - int(Player.Size.X)/2), Y: float32(rl.GetScreenHeight() - int(Player.Size.Y))}
	Player.Speed = 400
	Player.Lives = 5

	Ball.Radius = 20
	Ball.IsActive = false
	Ball.Speed = rl.Vector2{X: 0, Y: 0}
	Ball.Pos = rl.Vector2{X: Player.Pos.X + (Player.Size.X / 2), Y: Player.Pos.Y}

	for i := 0; i < BrickRows; i++ {
		for j := 0; j < BrickColumns; j++ {
			Bricks[i][j] = structs.Brick{}
			Bricks[i][j].Color = int(rl.GetRandomValue(0, 1))
			Bricks[i][j].IsActive = true
			Bricks[i][j].Size = rl.Vector2{X: float32(rl.GetScreenWidth() / 20), Y: float32(rl.GetScreenHeight() / 20)}
			Bricks[i][j].Pos = rl.Vector2{X: float32((rl.GetScreenWidth())/20 + (rl.GetScreenWidth() / 20 * (j - 1)) + (j * Offset) + rl.GetScreenWidth()/10), Y: float32((rl.GetScreenHeight())/20 + (rl.GetScreenHeight() / 20 * (i - 1)) + i*Offset)}
		}
	}

}

func GamePlayUpdate() {
	GamePlayDraw()
	GamePlayInput()
	rl.PlayMusicStream(Music)
	rl.UpdateMusicStream(Music)

	if !win {
		if !pause {
			if Ball.IsActive {
				Ball.Pos.X += Ball.Speed.X * rl.GetFrameTime()
				Ball.Pos.Y += Ball.Speed.Y * rl.GetFrameTime()
			} else {
				Ball.Pos = rl.Vector2{X: Player.Pos.X + Player.Size.X/2, Y: Player.Pos.Y - Ball.Radius}
			}

			if Ball.Pos.X-Ball.Radius < 0 {
				if Ball.Speed.X < 0 {
					Ball.Speed.X *= -1
				}
			}

			if Ball.Pos.X+Ball.Radius > float32(rl.GetScreenWidth()) {
				if Ball.Speed.X > 0 {
					Ball.Speed.X *= -1

				}
			}

			if Ball.Pos.Y-Ball.Radius <= 0 {
				if Ball.Speed.Y < 0 {
					Ball.Speed.Y *= -1
				}
			}

			if Ball.Pos.Y+Ball.Radius >= float32(rl.GetScreenHeight()) {
				Ball.Speed = rl.Vector2{X: 0, Y: 0}
				Ball.Pos = rl.Vector2{X: Player.Pos.X + Player.Size.X/2, Y: Player.Pos.Y - Ball.Radius}
				Player.Lives--
				Ball.IsActive = !Ball.IsActive
			}

			if rl.CheckCollisionCircleRec(Ball.Pos, Ball.Radius, Player.GetRect()) {
				if Ball.Speed.Y > 0 {
					Ball.Speed.Y *= -1
					Ball.Speed.X = (Ball.Pos.X - Player.Pos.X) / (Player.Size.X / 2) * Ball.Speed.Y
					if Ball.IsActive {
						rl.PlaySound(Poing)
					}
				}
			}

			for i := 0; i < BrickRows; i++ {
				for j := 0; j < BrickColumns; j++ {
					if Bricks[i][j].IsActive {
						if rl.CheckCollisionCircleRec(Ball.Pos, Ball.Radius, Bricks[i][j].GetRect()) {

						}
					}
				}
			}
		}
	}

}

func GamePlayDraw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	isColorIndexSet := false

	if !win {
		if !pause {
			rl.DrawRectangle(int32(Player.Pos.X), int32(Player.Pos.Y), int32(Player.Size.X), int32(Player.Size.Y), rl.Blue)
			rl.DrawCircle(int32(Ball.Pos.X), int32(Ball.Pos.Y), Ball.Radius, rl.Red)

			for i := 0; i < BrickRows; i++ {
				for j := 0; j < BrickColumns; j++ {
					colorIndex := 0
					if isColorIndexSet {
						colorIndex = 1
					}
					isColorIndexSet = !isColorIndexSet
					brick := Bricks[i][j]
					if brick.IsActive {
						rl.DrawRectangle(int32(brick.Pos.X), int32(brick.Pos.Y), int32(brick.Size.X), int32(brick.Size.Y), Color[colorIndex])
					}
				}
			}
		} else {
			rl.DrawText("Press P to play again", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/9, 40, rl.Red)
			rl.DrawText("Press enter to go to menu", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/6, 40, rl.Red)
		}
	} else {
		if Points < 30 {
			rl.DrawText("Game over", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/16, 40, rl.Red)
			rl.DrawText("Press P to play again", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/8, 40, rl.SkyBlue)
			rl.DrawText("Press enter to go to menu", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/5, 40, rl.SkyBlue)
		} else {
			rl.DrawText("You win!", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/16, 40, rl.Red)
			rl.DrawText("Press P to play again", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/8, 40, rl.SkyBlue)
			rl.DrawText("Press enter to go to menu", int32(rl.GetScreenWidth())/9, int32(rl.GetScreenHeight())/5, 40, rl.SkyBlue)
		}
	}

	rl.EndDrawing()
}

func GamePlayInput() {
	if !win {
		if rl.IsKeyPressed(rl.KeyP) {
			pause = !pause
		}

		if !pause {

			if Player.Pos.X > 0 {
				if rl.IsKeyDown(rl.KeyLeft) {
					Player.Pos.X -= float32(Player.Speed) * rl.GetFrameTime()
					if !Ball.IsActive {
						Ball.Pos.X -= float32(Player.Speed) * rl.GetFrameTime()
					}
				}
			}

			if (Player.Pos.X + Player.Size.X) < float32(rl.GetScreenWidth()) {
				if rl.IsKeyDown(rl.KeyRight) {
					Player.Pos.X += float32(Player.Speed) * rl.GetFrameTime()
					if !Ball.IsActive {
						Ball.Pos.X -= float32(Player.Speed) * rl.GetFrameTime()
					}
				}
			}

			if !Ball.IsActive {
				if rl.IsKeyPressed(rl.KeySpace) {
					randVal := rl.GetRandomValue(0, 1)
					Ball.IsActive = !Ball.IsActive
					if randVal == 0 {
						Ball.Speed = rl.Vector2{X: -100, Y: -300}
					} else {
						Ball.Speed = rl.Vector2{X: 100, Y: -300}
					}
				}
			}
		} else {
			if rl.IsKeyPressed(rl.KeyEnter) {
				pause = false
				screenType.CurrentScreenType = screenType.ScreenType.Menu
				GameInit()
			}
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			screenType.CurrentScreenType = screenType.ScreenType.Menu
			pause = false
			win = false
			GameInit()
		}

		if rl.IsKeyPressed(rl.KeyP) {
			pause = false
			win = false
			GameInit()
		}
	}
}
