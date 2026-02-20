package main

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - cubic-bezier lines")

	start := vector2.NewFloat32(0, 0)
	end := vector2.NewFloat32(float32(screenWidth), float32(screenHeight))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			start = rl.GetMousePosition()
		} else if rl.IsMouseButtonDown(rl.MouseRightButton) {
			end = rl.GetMousePosition()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("USE MOUSE LEFT-RIGHT CLICK to DEFINE LINE START and END POINTS", 15, 20, 20, rl.Gray)

		rl.DrawLineBezier(start, end, 2.0, rl.Red)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
