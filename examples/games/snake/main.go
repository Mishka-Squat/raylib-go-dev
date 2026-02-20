package main

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

const (
	snakeLength = 256
	squareSize  = 31
)

// Snake type
type Snake struct {
	Position vector2.Float32
	Size     vector2.Float32
	Speed    vector2.Float32
	Color    colorex.RGBA
}

// Food type
type Food struct {
	Position vector2.Float32
	Size     vector2.Float32
	Active   bool
	Color    colorex.RGBA
}

// Game type
type Game struct {
	ScreenWidth  int32
	ScreenHeight int32

	FramesCounter int32
	GameOver      bool
	Pause         bool

	Fruit         Food
	Snake         []Snake
	SnakePosition []vector2.Float32
	AllowMove     bool
	Offset        vector2.Float32
	CounterTail   int
}

func main() {
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, "sample game: snake")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game.Update()

		game.Draw()
	}

	rl.CloseWindow()
}

// Init - Initialize game
func (g *Game) Init() {
	g.ScreenWidth = 800
	g.ScreenHeight = 450

	g.FramesCounter = 0
	g.GameOver = false
	g.Pause = false

	g.CounterTail = 1
	g.AllowMove = false

	g.Offset = vector2.Float32{}
	g.Offset.X = float32(g.ScreenWidth % squareSize)
	g.Offset.Y = float32(g.ScreenHeight % squareSize)

	g.Snake = make([]Snake, snakeLength)

	for i := 0; i < snakeLength; i++ {
		g.Snake[i].Position = vector2.NewFloat32(g.Offset.X/2, g.Offset.Y/2)
		g.Snake[i].Size = vector2.NewFloat32(squareSize, squareSize)
		g.Snake[i].Speed = vector2.NewFloat32(squareSize, 0)

		if i == 0 {
			g.Snake[i].Color = rl.DarkBlue
		} else {
			g.Snake[i].Color = rl.Blue
		}
	}

	g.SnakePosition = make([]vector2.Float32, snakeLength)

	for i := 0; i < snakeLength; i++ {
		g.SnakePosition[i] = vector2.NewFloat32(0.0, 0.0)
	}

	g.Fruit.Size = vector2.NewFloat32(squareSize, squareSize)
	g.Fruit.Color = rl.SkyBlue
	g.Fruit.Active = false
}

// Update - Update game
func (g *Game) Update() {
	if !g.GameOver {
		if rl.IsKeyPressed(rl.KeyP) {
			g.Pause = !g.Pause
		}

		if !g.Pause {
			// control
			if rl.IsKeyPressed(rl.KeyRight) && g.Snake[0].Speed.X == 0 && g.AllowMove {
				g.Snake[0].Speed = vector2.NewFloat32(squareSize, 0)
				g.AllowMove = false
			}
			if rl.IsKeyPressed(rl.KeyLeft) && g.Snake[0].Speed.X == 0 && g.AllowMove {
				g.Snake[0].Speed = vector2.NewFloat32(-squareSize, 0)
				g.AllowMove = false
			}
			if rl.IsKeyPressed(rl.KeyUp) && g.Snake[0].Speed.Y == 0 && g.AllowMove {
				g.Snake[0].Speed = vector2.NewFloat32(0, -squareSize)
				g.AllowMove = false
			}
			if rl.IsKeyPressed(rl.KeyDown) && g.Snake[0].Speed.Y == 0 && g.AllowMove {
				g.Snake[0].Speed = vector2.NewFloat32(0, squareSize)
				g.AllowMove = false
			}

			// movement
			for i := 0; i < g.CounterTail; i++ {
				g.SnakePosition[i] = g.Snake[i].Position
			}

			if g.FramesCounter%5 == 0 {
				for i := 0; i < g.CounterTail; i++ {
					if i == 0 {
						g.Snake[0].Position.X += g.Snake[0].Speed.X
						g.Snake[0].Position.Y += g.Snake[0].Speed.Y
						g.AllowMove = true
					} else {
						g.Snake[i].Position = g.SnakePosition[i-1]
					}
				}
			}

			// wall behaviour
			if ((g.Snake[0].Position.X) > (float32(g.ScreenWidth) - g.Offset.X)) ||
				((g.Snake[0].Position.Y) > (float32(g.ScreenHeight) - g.Offset.Y)) ||
				(g.Snake[0].Position.X < 0) || (g.Snake[0].Position.Y < 0) {
				g.GameOver = true
			}

			// collision with yourself
			for i := 1; i < g.CounterTail; i++ {
				if (g.Snake[0].Position.X == g.Snake[i].Position.X) && (g.Snake[0].Position.Y == g.Snake[i].Position.Y) {
					g.GameOver = true
				}
			}

			if !g.Fruit.Active {
				g.Fruit.Active = true
				g.Fruit.Position = vector2.NewFloat32(
					float32(rl.GetRandomValue(0, (g.ScreenWidth/squareSize)-1)*squareSize)+(g.Offset.X)/2,
					float32(rl.GetRandomValue(0, (g.ScreenHeight/squareSize)-1)*squareSize)+(g.Offset.Y)/2,
				)

				for i := 0; i < g.CounterTail; i++ {
					for (g.Fruit.Position.X == g.Snake[i].Position.X) && (g.Fruit.Position.Y == g.Snake[i].Position.Y) {
						g.Fruit.Position = vector2.NewFloat32(
							float32(rl.GetRandomValue(0, (g.ScreenWidth/squareSize)-1)*squareSize)+g.Offset.X/2,
							float32(rl.GetRandomValue(0, (g.ScreenHeight/squareSize)-1)*squareSize)+g.Offset.Y/2,
						)
						i = 0
					}
				}
			}

			// collision
			if rl.CheckCollisionRecs(
				rl.NewRectangle(g.Snake[0].Position.X, g.Snake[0].Position.Y, g.Snake[0].Size.X, g.Snake[0].Size.Y),
				rl.NewRectangle(g.Fruit.Position.X, g.Fruit.Position.Y, g.Fruit.Size.X, g.Fruit.Size.Y),
			) {
				g.Snake[g.CounterTail].Position = g.SnakePosition[g.CounterTail-1]
				g.CounterTail += 1
				g.Fruit.Active = false
			}

			g.FramesCounter++
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.Init()
			g.GameOver = false
		}
	}
}

// Draw - Draw game
func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	if !g.GameOver {
		// Draw grid lines
		for i := int32(0); i < g.ScreenWidth/squareSize+1; i++ {
			rl.DrawLineV(
				vector2.NewFloat32(float32(squareSize*i)+g.Offset.X/2, g.Offset.Y/2),
				vector2.NewFloat32(float32(squareSize*i)+g.Offset.X/2, float32(g.ScreenHeight)-g.Offset.Y/2),
				rl.LightGray,
			)
		}

		for i := int32(0); i < g.ScreenHeight/squareSize+1; i++ {
			rl.DrawLineV(
				vector2.NewFloat32(g.Offset.X/2, float32(squareSize*i)+g.Offset.Y/2),
				vector2.NewFloat32(float32(g.ScreenWidth)-g.Offset.X/2, float32(squareSize*i)+g.Offset.Y/2),
				rl.LightGray,
			)
		}

		// Draw snake
		for i := 0; i < g.CounterTail; i++ {
			rl.DrawRectangleV(g.Snake[i].Position, g.Snake[i].Size, g.Snake[i].Color)
		}

		// Draw fruit to pick
		rl.DrawRectangleV(g.Fruit.Position, g.Fruit.Size, g.Fruit.Color)

		if g.Pause {
			rl.DrawText("GAME PAUSED", g.ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, g.ScreenHeight/2-40, 40, rl.Gray)
		}
	} else {
		rl.DrawText("PRESS [ENTER] TO PLAY AGAIN", int32(rl.GetScreenWidth())/2-rl.MeasureText("PRESS [ENTER] TO PLAY AGAIN", 20)/2, int32(rl.GetScreenHeight())/2-50, 20, rl.Gray)
	}

	rl.EndDrawing()
}
