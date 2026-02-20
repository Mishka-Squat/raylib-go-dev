package main

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - cubesmap loading and drawing")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(16.0, 14.0, 16.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 45.0

	image := rl.LoadImage("cubicmap.png")      // Load cubicmap image (RAM)
	cubicmap := rl.LoadTextureFromImage(image) // Convert image to texture to display (VRAM)

	mesh := rl.GenMeshCubicmap(image, vector3.NewFloat32(1.0, 1.0, 1.0))
	model := rl.LoadModelFromMesh(mesh)

	// NOTE: By default each cube is mapped to one part of texture atlas
	texture := rl.LoadTexture("cubicmap_atlas.png")                // Load map texture
	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture) // Set map diffuse texture

	mapPosition := vector3.NewFloat32(-16.0, 0.0, -8.0) // Set model position

	rl.UnloadImage(&image) // Unload cubicmap image from RAM, already uploaded to VRAM

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update

		rl.UpdateCamera(&camera, rl.CameraOrbital) // Update camera with orbital camera mode

		// Draw

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawModel(model, mapPosition, 1.0, rl.White)

		rl.EndMode3D()

		rl.DrawTextureEx(cubicmap, vector2.NewFloat32(float32(screenWidth-cubicmap.Width*4-20), 20), 0.0, 4.0, rl.White)
		rl.DrawRectangleLines(screenWidth-cubicmap.Width*4-20, 20, cubicmap.Width*4, cubicmap.Height*4, rl.Green)

		rl.DrawText("cubicmap image used to", 658, 90, 10, rl.Gray)
		rl.DrawText("generate map 3d model", 658, 104, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadTexture(&cubicmap) // Unload cubicmap texture
	rl.UnloadTexture(&texture)  // Unload map texture
	rl.UnloadModel(&model)      // Unload map model

	rl.CloseWindow()
}
