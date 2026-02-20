package main

import (
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - color selection (collision detection)")

	colors := [21]colorex.RGBA{
		rl.DarkGray, rl.Maroon, rl.Orange, rl.DarkGreen, rl.DarkBlue, rl.DarkPurple,
		rl.DarkBrown, rl.Gray, rl.Red, rl.Gold, rl.Lime, rl.Blue, rl.Violet, rl.Brown,
		rl.LightGray, rl.Pink, rl.Yellow, rl.Green, rl.SkyBlue, rl.Purple, rl.Beige,
	}

	colorsRecs := make([]rect2.Float32, 21) // Rectangles array

	// Fills colorsRecs data (for every rectangle)
	for i := 0; i < 21; i++ {
		r := rect2.Float32{}
		r.Position.X = float32(20 + 100*(i%7) + 10*(i%7))
		r.Position.Y = float32(60 + 100*(i/7) + 10*(i/7))
		r.Size.X = 100
		r.Size.Y = 100

		colorsRecs[i] = r
	}

	selected := make([]bool, 21) // Selected rectangles indicator

	mousePoint := vector2.Float32{}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		mousePoint = rl.GetMousePosition()

		for i := 0; i < 21; i++ { // Iterate along all the rectangles
			if rl.CheckCollisionPointRec(mousePoint, colorsRecs[i]) {
				colors[i].A = 120

				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					selected[i] = !selected[i]
				}
			} else {
				colors[i].A = 255
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := 0; i < 21; i++ { // Draw all rectangles
			rl.DrawRectangleRec(colorsRecs[i], colors[i])

			// Draw four rectangles around selected rectangle
			if selected[i] {
				rl.DrawRectangle(int32(colorsRecs[i].Position.X), int32(colorsRecs[i].Position.Y), 100, 10, rl.RayWhite)    // Square top rectangle
				rl.DrawRectangle(int32(colorsRecs[i].Position.X), int32(colorsRecs[i].Position.Y), 10, 100, rl.RayWhite)    // Square left rectangle
				rl.DrawRectangle(int32(colorsRecs[i].Position.X+90), int32(colorsRecs[i].Position.Y), 10, 100, rl.RayWhite) // Square right rectangle
				rl.DrawRectangle(int32(colorsRecs[i].Position.X), int32(colorsRecs[i].Position.Y)+90, 100, 10, rl.RayWhite) // Square bottom rectangle
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
