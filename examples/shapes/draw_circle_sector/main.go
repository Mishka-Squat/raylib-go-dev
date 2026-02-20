/*******************************************************************************************
 *
 *   raylib [shapes] example - draw circle sector (with gui options)
 *
 *   Example originally created with raylib 2.5, last time updated with raylib 2.5
 *
 *   Example contributed by Vlad Adrian (@demizdor) and reviewed by Ramon Santamaria (@raysan5)
 *
 *   Example licensed under an unmodified zlib/libpng license, which is an OSI-certified,
 *   BSD-like license that allows static linking with closed source software
 *
 *   Copyright (c) 2018-2024 Vlad Adrian (@demizdor) and Ramon Santamaria (@raysan5)
 *
 ********************************************************************************************/
package main

import (
	"fmt"
	"math"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	gui "github.com/Mishka-Squat/raylib-go/raygui"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - draw circle sector")

	center := vector2.Float32{X: (float32(screenWidth) - 300) / 2.0, Y: float32(screenHeight / 2.0)}
	var outerRadius, startAngle, endAngle, segments, minSegments float32 = 180.0, 0.0, 180.0, 10.0, 4

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawLine(500, 0, 500, screenWidth, rl.Fade(rl.LightGray, 0.6))
		rl.DrawRectangle(500, 0, screenWidth-500, screenHeight, rl.Fade(rl.LightGray, 0.3))

		rl.DrawCircleSector(center, outerRadius, startAngle, endAngle, int32(segments), rl.Fade(rl.Maroon, 0.3))
		rl.DrawCircleSectorLines(center, outerRadius, startAngle, endAngle, int32(segments), rl.Fade(rl.Maroon, 0.6))

		// Draw GUI controls
		r := rect2.NewFloat32(vector2.NewFloat32(600, 40), vector2.NewFloat32(120, 20))
		msg := fmt.Sprintf("%.2f", startAngle)
		startAngle = gui.Slider(r, "StartAngle", msg, startAngle, 0, 720)

		r = rect2.NewFloat32(vector2.NewFloat32(600, 70), vector2.NewFloat32(120, 20))
		msg = fmt.Sprintf("%.2f", endAngle)
		endAngle = gui.Slider(r, "EndAngle", msg, endAngle, 0, 720)

		r = rect2.NewFloat32(vector2.NewFloat32(600, 140), vector2.NewFloat32(120, 20))
		msg = fmt.Sprintf("%.2f", outerRadius)
		outerRadius = gui.Slider(r, "Radius", msg, outerRadius, 0, 200)

		r = rect2.NewFloat32(vector2.NewFloat32(600, 170), vector2.NewFloat32(120, 20))
		msg = fmt.Sprintf("%.2f", segments)
		segments = gui.Slider(r, "Segments", msg, segments, 0, 100)

		minSegments = calculateMinSegments(startAngle, endAngle)
		text := "MODE: AUTO"
		color := rl.DarkGray
		if segments >= minSegments {
			text = "MODE: MANUAL"
			color = rl.Maroon
		}
		rl.DrawText(text, 600, 200, 10, color)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	// De-Initialization
	rl.CloseWindow() // Close window and OpenGL context
}

func calculateMinSegments(startAngle, endAngle float32) float32 {
	return float32(math.Trunc(math.Ceil((float64(endAngle - startAngle)) / 90)))
}
