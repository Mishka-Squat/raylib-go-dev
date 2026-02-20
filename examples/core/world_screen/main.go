package main

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera free")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(10.0, 10.0, 10.0) // Camera position
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)      // Camera looking at point
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)          // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                                     // Camera field-of-view Y

	cubePosition := vector3.NewFloat32(0.0, 0.0, 0.0)
	cubeScreenPosition := vector2.Float32{}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree) // Update camera with free camera mode

		// Calculate cube screen space position (with a little offset to be in top)
		cubeScreenPosition = rl.GetWorldToScreen(vector3.NewFloat32(cubePosition.X, cubePosition.Y+2.5, cubePosition.Z), camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawText("Enemy: 100 / 100", int32(cubeScreenPosition.X)-rl.MeasureText("Enemy: 100 / 100", 20)/2, int32(cubeScreenPosition.Y), 20, rl.Black)
		rl.DrawText("Text is always on top of the cube", (screenWidth-rl.MeasureText("Text is always on top of the cube", 20))/2, 25, 20, rl.Gray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
