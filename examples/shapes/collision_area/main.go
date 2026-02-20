package main

import (
	"fmt"

	"github.com/Mishka-Squat/gamemath/rect2"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

var (
	screenWidth  = int32(1280)
	screenHeight = int32(720)

	pause     = false
	collision = false

	boxCollision     = rect2.Float32{}
	screenUpperLimit = float32(40)
	boxAspeedX       = float32(4)
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - collision area")

	rl.SetTargetFPS(60)

	boxA := rl.MakeRectangle(10, float32(rl.GetScreenHeight()/2)-50, 200, 100)
	boxB := rl.MakeRectangle(float32(rl.GetScreenWidth()/2)-30, float32(rl.GetScreenHeight()/2)-30, 60, 60)

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()

		if !pause {
			boxA.Position.X += boxAspeedX
		}
		if boxA.Position.X+boxA.Size.X >= float32(rl.GetScreenWidth()) || boxA.Position.X <= 0 {
			boxAspeedX *= -1
		}

		boxB.Position.X = mousePos.X - boxB.Size.X/2
		boxB.Position.Y = mousePos.Y - boxB.Size.Y/2

		if boxB.Position.X+boxB.Size.X >= float32(rl.GetScreenWidth()) {
			boxB.Position.X = float32(rl.GetScreenWidth()) - boxB.Size.X
		} else if boxB.Position.X <= 0 {
			boxB.Position.X = 0
		}

		if boxB.Position.Y+boxB.Size.Y >= float32(rl.GetScreenHeight()) {
			boxB.Position.Y = float32(rl.GetScreenHeight()) - boxB.Size.Y
		} else if boxB.Position.X <= screenUpperLimit {
			boxB.Position.Y = screenUpperLimit
		}

		collision := rl.CheckCollisionRecs(boxA, boxB)

		if collision {
			boxCollision = rl.GetCollisionRec(boxA, boxB)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if collision {
			rl.DrawRectangle(0, 0, screenWidth, int32(screenUpperLimit), rl.Red)
			rl.DrawRectangleRec(boxCollision, rl.Lime)
			rl.DrawText("COLLISION", int32(rl.GetScreenWidth()/2)-(rl.MeasureText("COLLISION", 20)/2), int32(screenUpperLimit/2)-10, 20, rl.Black)
			rl.DrawText("Collision Area: "+fmt.Sprint(boxCollision.Size.X*boxCollision.Size.Y), int32(rl.GetScreenWidth()/2)-100, int32(screenUpperLimit+10), 20, rl.Black)
		} else {
			rl.DrawRectangle(0, 0, screenWidth, int32(screenUpperLimit), rl.Black)
		}

		rl.DrawRectangleRec(boxA, rl.Orange)
		rl.DrawRectangleRec(boxB, rl.Blue)

		rl.DrawText("Press SPACE to PAUSE/RESUME", 20, int32(rl.GetScreenHeight())-35, 20, rl.Black)

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
