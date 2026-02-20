package main

import (
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	ez "github.com/Mishka-Squat/raylib-go/easings"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

var (
	playTime = float32(240)
	fps      = int32(60)
)

const (
	recsW, recsH = 50, 50
	screenWidth  = int32(800)
	screenHeight = int32(450)
	maxRecsX     = int(screenWidth) / recsW
	maxRecsY     = int(screenHeight) / recsH
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - easings rectangle array")

	var recs [maxRecsX * maxRecsY]rect2.Float32

	for y := 0; y < maxRecsY; y++ {

		for x := 0; x < maxRecsX; x++ {
			recs[y*maxRecsX+x].Position.X = float32(recsW/2 + recsW*float32(x))
			recs[y*maxRecsX+x].Position.Y = float32(recsH/2 + recsH*float32(y))
			recs[y*maxRecsX+x].Size.X = float32(recsW)
			recs[y*maxRecsX+x].Size.Y = float32(recsH)
		}

	}

	rotation := float32(0)
	frameCount := float32(0)
	state := 0

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {

		if state == 0 {
			frameCount++

			for i := 0; i < maxRecsX*maxRecsY; i++ {
				recs[i].Size.Y = float32(ez.LinearIn(frameCount, recsH, -recsH, playTime))
				recs[i].Size.X = float32(ez.LinearIn(frameCount, recsW, -recsW, playTime))

				if recs[i].Size.Y < 0 {
					recs[i].Size.Y = 0
				}

				if recs[i].Size.X < 0 {
					recs[i].Size.X = 0
				}

				if recs[i].Size.Y == 0 && recs[i].Size.X == 0 {
					state = 1
				}
				rotation = float32(ez.LinearIn(frameCount, 0, 360, playTime))
			}
		} else if state == 1 && rl.IsKeyPressed(rl.KeySpace) {
			frameCount = 0
			for i := 0; i < maxRecsX*maxRecsY; i++ {
				recs[i].Size.Y = float32(recsH)
				recs[i].Size.X = float32(recsW)
			}

			state = 0
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if state == 0 {
			for i := 0; i < maxRecsX*maxRecsY; i++ {
				rl.DrawRectanglePro(recs[i], vector2.NewFloat32(recs[i].Size.X/2, recs[i].Size.Y/2), rotation, rl.Red)
			}
		} else if state == 1 {
			txtlen := rl.MeasureText("SPACE to replay", 20)
			rl.DrawText("SPACE to replay", (screenWidth/2)-txtlen/2, 200, 20, rl.Black)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
