package main

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

var (
	mouseScaleMarkSize              = float32(12)
	rec                             = rl.MakeRectangle(100, 100, 200, 80)
	mouseScaleReady, MouseScaleMode = false, false
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - rectangle scaling mouse")

	rl.SetTargetFPS(60)

	rl.SetMousePosition(0, 0)

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()

		if rl.CheckCollisionPointRec(mousePos, rl.MakeRectangle(rec.Position.X+rec.Size.X-mouseScaleMarkSize, rec.Position.Y+rec.Size.Y-mouseScaleMarkSize, mouseScaleMarkSize, mouseScaleMarkSize)) {
			mouseScaleReady = true
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				MouseScaleMode = true
			}
		} else {
			mouseScaleReady = false
		}

		if MouseScaleMode {

			mouseScaleReady = true
			rec.Size.X = mousePos.X - rec.Position.X
			rec.Size.Y = mousePos.Y - rec.Position.Y

			// CHECK MIN MAX REC SIZES
			if rec.Size.X < mouseScaleMarkSize {
				rec.Size.X = rec.Size.Y
			}
			if rec.Size.Y < mouseScaleMarkSize {
				rec.Size.Y = rec.Size.X
			}
			if rec.Size.X > (float32(rl.GetScreenWidth()) - rec.Position.X) {
				rec.Size.X = float32(rl.GetScreenWidth()) - rec.Position.X
			}
			if rec.Size.Y > (float32(rl.GetScreenHeight()) - rec.Position.Y) {
				rec.Size.Y = float32(rl.GetScreenHeight()) - rec.Position.Y
			}
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				MouseScaleMode = false
			}

		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Scale rectangle dragging from bottom-right corner!", 10, 10, 20, rl.Black)

		rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.5))

		if mouseScaleReady {
			rl.DrawRectangleLinesEx(rec, 1, rl.Red)
			rl.DrawTriangle(vector2.NewFloat32(rec.Position.X+rec.Size.X-mouseScaleMarkSize, rec.Position.Y+rec.Size.Y), vector2.NewFloat32(rec.Position.X+rec.Size.X, rec.Position.Y+rec.Size.Y), vector2.NewFloat32(rec.Position.X+rec.Size.X, rec.Position.Y+rec.Size.Y-mouseScaleMarkSize), rl.Red)
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
