package main

import (
	"fmt"

	"github.com/Mishka-Squat/gamemath/vector2"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

var (
	screenW    = int32(800)
	screenH    = int32(440)
	playerSize = float32(40)
	cam1, cam2 rl.Camera2D
)

func main() {

	rl.InitWindow(screenW, screenH, "raylib [core] example - 2d camera split screen")

	player1 := rl.NewRectangle(200, 200, playerSize, playerSize)
	player2 := rl.NewRectangle(250, 200, playerSize, playerSize)

	cam1.Target = vector2.NewFloat32(player1.Position.X, player1.Position.Y)
	cam1.Offset = vector2.NewFloat32(200, 200)
	cam1.Rotation = 0
	cam1.Zoom = 1

	cam2 = cam1
	cam2.Target = vector2.NewFloat32(player2.Position.X, player2.Position.Y)

	screenCam1 := rl.LoadRenderTexture(screenW/2, screenH)
	screenCam2 := rl.LoadRenderTexture(screenW/2, screenH)

	splitScreenRec := rl.NewRectangle(0, 0, float32(screenCam1.Texture.Width), -float32(screenCam1.Texture.Height))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyS) {
			player1.Position.Y += 3
		} else if rl.IsKeyDown(rl.KeyW) {
			player1.Position.Y -= 3
		}
		if rl.IsKeyDown(rl.KeyD) {
			player1.Position.X += 3
		} else if rl.IsKeyDown(rl.KeyA) {
			player1.Position.X -= 3
		}

		if rl.IsKeyDown(rl.KeyUp) {
			player2.Position.Y -= 3
		} else if rl.IsKeyDown(rl.KeyDown) {
			player2.Position.Y += 3
		}
		if rl.IsKeyDown(rl.KeyRight) {
			player2.Position.X += 3
		} else if rl.IsKeyDown(rl.KeyLeft) {
			player2.Position.X -= 3
		}

		cam1.Target = vector2.NewFloat32(player1.Position.X, player1.Position.Y)
		cam2.Target = vector2.NewFloat32(player2.Position.X, player2.Position.Y)

		rl.BeginTextureMode(screenCam1)
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(cam1)

		for i := 0; i < int(screenW/int32(playerSize))+1; i++ {
			rl.DrawLineV(vector2.NewFloat32(playerSize*float32(i), 0), vector2.NewFloat32(playerSize*float32(i), float32(screenH)), rl.LightGray)
		}
		for i := 0; i < int(screenH/int32(playerSize))+1; i++ {
			rl.DrawLineV(vector2.NewFloat32(0, playerSize*float32(i)), vector2.NewFloat32(float32(screenW), playerSize*float32(i)), rl.LightGray)
		}
		for i := 0; i < int(screenW/int32(playerSize)); i++ {
			for j := 0; j < int(screenH/int32(playerSize)); j++ {
				rl.DrawText("["+fmt.Sprint(i)+","+fmt.Sprint(j)+"]", 10+int32(playerSize*float32(i)), 15+int32(playerSize*float32(j)), 10, rl.LightGray)
			}
		}

		rl.DrawRectangleRec(player1, rl.Red)
		rl.DrawRectangleRec(player2, rl.Blue)
		rl.EndMode2D()

		rl.DrawRectangle(0, 0, screenW/2, 30, rl.Fade(rl.RayWhite, 0.6))
		rl.DrawText("PLAYER 1 WASD KEYS", 10, 10, 10, rl.Maroon)
		rl.EndTextureMode()

		rl.BeginTextureMode(screenCam2)
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(cam2)

		for i := 0; i < int(screenW/int32(playerSize))+1; i++ {
			rl.DrawLineV(vector2.NewFloat32(playerSize*float32(i), 0), vector2.NewFloat32(playerSize*float32(i), float32(screenH)), rl.LightGray)
		}
		for i := 0; i < int(screenH/int32(playerSize))+1; i++ {
			rl.DrawLineV(vector2.NewFloat32(0, playerSize*float32(i)), vector2.NewFloat32(float32(screenW), playerSize*float32(i)), rl.LightGray)
		}
		for i := 0; i < int(screenW/int32(playerSize)); i++ {
			for j := 0; j < int(screenH/int32(playerSize)); j++ {
				rl.DrawText("["+fmt.Sprint(i)+","+fmt.Sprint(j)+"]", 10+int32(playerSize*float32(i)), 15+int32(playerSize*float32(j)), 10, rl.LightGray)
			}
		}

		rl.DrawRectangleRec(player1, rl.Red)
		rl.DrawRectangleRec(player2, rl.Blue)
		rl.EndMode2D()

		rl.DrawRectangle(0, 0, screenW/2, 30, rl.Fade(rl.RayWhite, 0.6))
		rl.DrawText("PLAYER 2 ARROW KEYS", 10, 10, 10, rl.Maroon)
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTextureRec(screenCam1.Texture, splitScreenRec, vector2.NewFloat32(0, 0), rl.White)
		rl.DrawTextureRec(screenCam2.Texture, splitScreenRec, vector2.NewFloat32(float32(screenW/2), 0), rl.White)
		rl.DrawRectangle((screenW/2)-2, 0, 4, screenH, rl.LightGray)

		rl.EndDrawing()

	}

	rl.UnloadRenderTexture(&screenCam1)
	rl.UnloadRenderTexture(&screenCam2)

	rl.CloseWindow()
}
