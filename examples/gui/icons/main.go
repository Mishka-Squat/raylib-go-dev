package main

import (
	rg "github.com/Mishka-Squat/raylib-go/raygui"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "raylib-go - icons example")
	defer rl.CloseWindow()

	rg.LoadIcons("default_icons_with_255.rgi", false)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rg.DrawIcon(rg.ICON_255, 100, 100, 8, rl.Gray)
		rl.EndDrawing()
	}
}
