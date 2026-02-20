package main

import (
	"github.com/Mishka-Squat/gamemath/vector3"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - geometric shapes")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(0.0, 10.0, 10.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 45.0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(vector3.NewFloat32(-4.0, 0.0, 2.0), 2.0, 5.0, 2.0, rl.Red)
		rl.DrawCubeWires(vector3.NewFloat32(-4.0, 0.0, 2.0), 2.0, 5.0, 2.0, rl.Gold)
		rl.DrawCubeWires(vector3.NewFloat32(-4.0, 0.0, -2.0), 3.0, 6.0, 2.0, rl.Maroon)

		rl.DrawSphere(vector3.NewFloat32(-1.0, 0.0, -2.0), 1.0, rl.Green)
		rl.DrawSphereWires(vector3.NewFloat32(1.0, 0.0, 2.0), 2.0, 16, 16, rl.Lime)

		rl.DrawCylinder(vector3.NewFloat32(4.0, 0.0, -2.0), 1.0, 2.0, 3.0, 4, rl.SkyBlue)
		rl.DrawCylinderWires(vector3.NewFloat32(4.0, 0.0, -2.0), 1.0, 2.0, 3.0, 4, rl.DarkBlue)
		rl.DrawCylinderWires(vector3.NewFloat32(4.5, -1.0, 2.0), 1.0, 1.0, 2.0, 6, rl.Brown)

		rl.DrawCylinder(vector3.NewFloat32(1.0, 0.0, -4.0), 0.0, 1.5, 3.0, 8, rl.Gold)
		rl.DrawCylinderWires(vector3.NewFloat32(1.0, 0.0, -4.0), 0.0, 1.5, 3.0, 8, rl.Pink)

		rl.DrawGrid(10, 1.0) // Draw a grid

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
