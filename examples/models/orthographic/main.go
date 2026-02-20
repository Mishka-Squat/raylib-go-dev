package main

import (
	"github.com/Mishka-Squat/gamemath/vector3"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {

	fovyPerspective := float32(45)
	widthOrthographic := float32(10)

	screenWidth := int32(1280)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - orthographic")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(0.0, 10.0, 10.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = fovyPerspective
	camera.Projection = rl.CameraPerspective

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraOrbital)

		if rl.IsKeyPressed(rl.KeySpace) {

			if camera.Projection == rl.CameraPerspective {
				camera.Fovy = widthOrthographic
				camera.Projection = rl.CameraOrthographic
			} else {
				camera.Fovy = fovyPerspective
				camera.Projection = rl.CameraPerspective
			}

		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(vector3.NewFloat32(-4, 0, 2), 2, 5, 2, rl.Red)
		rl.DrawCubeWires(vector3.NewFloat32(-4, 0, 2), 2, 5, 2, rl.Gold)
		rl.DrawCubeWires(vector3.NewFloat32(-4, 0, -2), 3, 6, 2, rl.Maroon)

		rl.DrawSphere(vector3.NewFloat32(-1, 0, -2), 1, rl.Green)
		rl.DrawSphereWires(vector3.NewFloat32(1, 0, 2), 2, 16, 16, rl.Lime)

		rl.DrawCylinder(vector3.NewFloat32(4, 0, -2), 1, 2, 3, 4, rl.SkyBlue)
		rl.DrawCylinderWires(vector3.NewFloat32(4, 0, -2), 1, 2, 3, 4, rl.DarkBlue)
		rl.DrawCylinderWires(vector3.NewFloat32(4.5, -1, 2), 1, 1, 2, 6, rl.Brown)

		rl.DrawCylinder(vector3.NewFloat32(1, 0, -4), 0, 1.5, 3, 8, rl.Gold)
		rl.DrawCylinderWires(vector3.NewFloat32(1, 0, -4), 0, 1.5, 3, 8, rl.Pink)

		rl.DrawGrid(10, 1)

		rl.EndMode3D()

		rl.DrawText("PRESS SPACE TO CHANGE CAMERA PROJECTION", 10, 10, 20, rl.Black)

		if camera.Projection == rl.CameraPerspective {
			rl.DrawText("CURRENT CAMERA PROJECTION IS PERSPECTIVE", 10, 30, 20, rl.Black)
		} else {
			rl.DrawText("CURRENT CAMERA PROJECTION IS ORTHOGRAPHIC", 10, 30, 20, rl.Black)
		}

		rl.DrawFPS(screenWidth-100, 10)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
