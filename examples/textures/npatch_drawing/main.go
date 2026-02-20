/*******************************************************************************************
*
*   raylib [textures] example - N-patch drawing
*
*   NOTE: Images are loaded in CPU memory (RAM); textures are loaded in GPU memory (VRAM)
*
*   Example originally created with raylib 2.0, last time updated with raylib 2.5
*
*   Example contributed by Jorge A. Gomes (@overdev) and reviewed by Ramon Santamaria (@raysan5)
*
*   Example licensed under an unmodified zlib/libpng license, which is an OSI-certified,
*   BSD-like license that allows static linking with closed source software
*
*   Copyright (c) 2018-2024 Jorge A. Gomes (@overdev) and Ramon Santamaria (@raysan5)
*
********************************************************************************************/
package main

import (
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	gm "github.com/Mishka-Squat/goex/mathex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] example - N-patch drawing")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)
	nPatchTexture := rl.LoadTexture("ninepatch_button.png")

	var mousePosition vector2.Float32
	var origin vector2.Float32

	// Position and size of the n-patches
	dstRec1 := rect2.NewFloat32(vector2.NewFloat32(480.0, 160.0), vector2.NewFloat32(32.0, 32.0))
	dstRec2 := rect2.NewFloat32(vector2.NewFloat32(160.0, 160.0), vector2.NewFloat32(32.0, 32.0))
	dstRecH := rect2.NewFloat32(vector2.NewFloat32(160.0, 93.0), vector2.NewFloat32(32.0, 32.0))
	dstRecV := rect2.NewFloat32(vector2.NewFloat32(92.0, 160.0), vector2.NewFloat32(32.0, 32.0))

	// A 9-patch (NPatchNinePatch) changes its sizes in both axis
	ninePatchInfo1 := rl.NPatchInfo{
		Source: rect2.NewFloat32(vector2.Zero[float32](), vector2.NewFloat32(64.0, 64.0)),
		Left:   12,
		Top:    40,
		Right:  12,
		Bottom: 12,
		Layout: rl.NPatchNinePatch,
	}
	ninePatchInfo2 := rl.NPatchInfo{
		Source: rect2.NewFloat32(vector2.NewFloat32(0, 128.0), vector2.NewFloat32(64.0, 64.0)),
		Left:   16,
		Top:    16,
		Right:  16,
		Bottom: 16,
		Layout: rl.NPatchNinePatch,
	}

	// A horizontal 3-patch (NPatchThreePatchHorizontal) changes its sizes along the x-axis only
	h3PatchInfo := rl.NPatchInfo{
		Source: rect2.NewFloat32(vector2.NewFloat32(0, 64.0), vector2.NewFloat32(64.0, 64.0)),
		Left:   8,
		Top:    8,
		Right:  8,
		Bottom: 8,
		Layout: rl.NPatchThreePatchHorizontal,
	}

	// A vertical 3-patch (NPatchThreePatchVertical) changes its sizes along the y-axis only
	v3PatchInfo := rl.NPatchInfo{
		Source: rect2.NewFloat32(vector2.NewFloat32(0, 192.0), vector2.NewFloat32(64.0, 64.0)),
		Left:   6,
		Top:    6,
		Right:  6,
		Bottom: 6,
		Layout: rl.NPatchThreePatchVertical,
	}

	rl.SetTargetFPS(60)

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		mousePosition = rl.GetMousePosition()

		// Resize the n-patches based on mouse position
		dstRec1.SetWidth(mousePosition.X - dstRec1.X())
		dstRec1.SetHeight(mousePosition.Y - dstRec1.Y())
		dstRec2.SetWidth(mousePosition.X - dstRec2.X())
		dstRec2.SetHeight(mousePosition.Y - dstRec2.Y())
		dstRecH.SetWidth(mousePosition.X - dstRecH.X())
		dstRecV.SetHeight(mousePosition.Y - dstRecV.Y())

		// Set a minimum Width and/or Height
		dstRec1.SetWidth(gm.Clamp(dstRec1.Width(), 1, 300))
		dstRec1.SetHeight(gm.Clamp(dstRec1.Height(), 1, screenHeight))
		dstRec2.SetWidth(gm.Clamp(dstRec2.Width(), 1, 300))
		dstRec2.SetHeight(gm.Clamp(dstRec2.Height(), 1, screenHeight))
		dstRecH.SetWidth(gm.Clamp(dstRecH.Width(), 1, screenWidth))
		dstRecV.SetHeight(gm.Clamp(dstRecV.Height(), 1, screenHeight))

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw the n-patches
		rl.DrawTextureNPatch(nPatchTexture, ninePatchInfo2, dstRec2, origin, 0.0, rl.White)
		rl.DrawTextureNPatch(nPatchTexture, ninePatchInfo1, dstRec1, origin, 0.0, rl.White)
		rl.DrawTextureNPatch(nPatchTexture, h3PatchInfo, dstRecH, origin, 0.0, rl.White)
		rl.DrawTextureNPatch(nPatchTexture, v3PatchInfo, dstRecV, origin, 0.0, rl.White)

		// Draw the source texture
		rl.DrawRectangleLines(5, 88, 74, 266, rl.Blue)
		rl.DrawTexture(nPatchTexture, 10, 93, rl.White)
		rl.DrawText("TEXTURE", 15, 360, 10, rl.DarkGray)

		rl.DrawText("Move the mouse to stretch or shrink the n-patches", 10, 20, 20, rl.DarkGray)

		rl.EndDrawing()
	}

	// De-Initialization
	rl.UnloadTexture(&nPatchTexture) // Texture unloading
	rl.CloseWindow()                 // Close window and OpenGL context
}
