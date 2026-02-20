package main

import (
	"fmt"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	// Enable config flags for resizable window and vertical synchro
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - Window Scale Letterbox")
	rl.SetWindowMinSize(320, 240)

	gameScreenWidth, gameScreenHeight := int32(640), int32(480)

	// Render texture initialization, used to hold the rendering result so we can easily resize it
	target := rl.LoadRenderTexture(gameScreenWidth, gameScreenHeight)
	rl.SetTextureFilter(target.Texture, rl.FilterBilinear) // Texture scale filter to use

	colors := getRandomColors()

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Compute required frame buffer scaling
		scale := min(float32(rl.GetScreenWidth())/float32(gameScreenWidth),
			float32(rl.GetScreenHeight())/float32(gameScreenHeight))

		if rl.IsKeyPressed(rl.KeySpace) {
			// Recalculate random colors for the bars
			colors = getRandomColors()
		}

		// Update virtual mouse (clamped mouse value behind game screen)
		mouse := rl.GetMousePosition()
		virtualMouse := vector2.Float32{
			X: (mouse.X - (float32(rl.GetScreenWidth())-(float32(gameScreenWidth)*scale))*0.5) / scale,
			Y: (mouse.Y - (float32(rl.GetScreenHeight())-(float32(gameScreenHeight)*scale))*0.5) / scale,
		}
		virtualMouse = virtualMouse.Clamp0V(vector2.NewFloat32(gameScreenWidth, gameScreenHeight))

		// Apply the same transformation as the virtual mouse to the real mouse (i.e. to work with raygui)
		//rl.SetMouseOffset(
		//	int(-(float32(rl.GetScreenWidth())-(float32(gameScreenWidth)*scale))*0.5),
		//	int(-(float32(rl.GetScreenHeight())-(float32(gameScreenHeight)*scale))*0.5),
		//)
		//rl.SetMouseScale(1/scale, 1/scale)

		// Draw everything in the render texture, note this will not be rendered on screen, yet
		rl.BeginTextureMode(target)

		rl.ClearBackground(rl.White)
		for i := int32(0); i < 10; i++ {
			rl.DrawRectangle(0, (gameScreenHeight/10)*i, gameScreenWidth, gameScreenHeight/10, colors[i])
		}

		text := "If executed inside a window,\nyou can resize the window,\nand see the screen scaling!"
		rl.DrawText(text, 10, 25, 20, rl.White)
		text = fmt.Sprintf("Default Mouse: [%.0f , %.0f]", mouse.X, mouse.Y)
		rl.DrawText(text, 350, 25, 20, rl.Green)
		text = fmt.Sprintf("Virtual Mouse: [%.0f , %.0f]", virtualMouse.X, virtualMouse.Y)
		rl.DrawText(text, 350, 55, 20, rl.Yellow)

		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// Draw render texture to screen, properly scaled
		rl.DrawTexturePro(
			target.Texture,
			rect2.NewFloat32(vector2.Zero[float32](), vector2.NewFloat32(target.Texture.Width, -target.Texture.Height)),
			rect2.NewFloat32(
				rl.GetScreenSize().ToFloat32().SubXY(float32(gameScreenWidth)*scale*0.5, float32(gameScreenHeight)*scale*0.5),
				vector2.NewFloat32(float32(gameScreenWidth)*scale, float32(gameScreenHeight)*scale),
			),
			vector2.Float32{X: 0, Y: 0}, 0, rl.White,
		)
		rl.EndDrawing()
	}
	rl.UnloadRenderTexture(&target)
	rl.CloseWindow() // Close window and OpenGL context
}

func getRandomColors() []colorex.RGBA {
	var colors []colorex.RGBA

	for i := 0; i < 10; i++ {
		randomColor := colorex.RGBA{
			R: rndUint8(100, 250),
			G: rndUint8(50, 150),
			B: rndUint8(10, 100),
			A: 255,
		}
		colors = append(colors, randomColor)
	}

	return colors
}

func rndUint8(min, max int32) uint8 {
	return uint8(rl.GetRandomValue(min, max))
}
