package main

import (
	"github.com/Mishka-Squat/gamemath/vector3"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - texture tiling")

	cam := rl.Camera3D{}
	cam.Position = vector3.NewFloat32(4, 4, 4)
	cam.Target = vector3.NewFloat32(0, 0.5, 0)
	cam.Up = vector3.NewFloat32(0, 1, 0)
	cam.Fovy = 45
	cam.Projection = rl.CameraPerspective

	cube := rl.GenMeshCube(1, 1, 1)
	model := rl.LoadModelFromMesh(cube)

	texture := rl.LoadTexture("cubicmap_atlas.png")
	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture)

	tiling := []float32{3, 3}
	shader := rl.LoadShader("", "tiling.fs")
	rl.SetShaderValueA(shader, rl.GetShaderLocation(shader, "tiling"), tiling, rl.ShaderUniformVec2)
	model.Materials.Shader = shader

	rl.DisableCursor()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&cam, rl.CameraOrbital)

		if rl.IsKeyPressed(rl.KeyZ) {
			cam.Target = vector3.NewFloat32(0, 0.5, 0)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(cam)

		rl.BeginShaderMode(shader)
		rl.DrawModel(model, vector3.Zero[float32](), 2, rl.White)
		rl.EndShaderMode()

		rl.DrawGrid(10, 1)

		rl.EndMode3D()

		rl.EndDrawing()
	}

	rl.UnloadShader(&shader)
	rl.UnloadModel(&model)
	rl.UnloadTexture(&texture)

	rl.CloseWindow()
}
