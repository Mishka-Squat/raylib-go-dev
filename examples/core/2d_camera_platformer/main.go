package main

import (
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

const (
	screenWidth     = 800
	screenHeight    = 450
	gravity         = 400
	playerJumpSpeed = 350.0
	playerHORSpeed  = 200.0
)

type Player struct {
	position vector2.Float32
	speed    float32
	canJump  bool
}

type EnvironmentItem struct {
	rect     rect2.Float32
	blocking bool
	color    colorex.RGBA
}

type cameraUpdater func(*rl.Camera2D, *Player, []EnvironmentItem, float32)

// These 3 variables are used only for camera 4,
// but they need to be declared on module level
// for camera 4 to work (static in C)
var eveningOut = false
var evenOutSpeed float32 = 700
var evenOutTarget float32 = 0

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 2d camera platformer")

	player := Player{
		position: vector2.NewFloat32(400, 280),
		speed:    0,
		canJump:  false,
	}

	envItems := []EnvironmentItem{
		{rect: rect2.NewFloat32(vector2.Zero[float32](), vector2.NewFloat32(1000, 400)), blocking: false, color: rl.LightGray},
		{rect: rect2.NewFloat32(vector2.NewFloat32(0, 400), vector2.NewFloat32(1000, 200)), blocking: true, color: rl.Gray},
		{rect: rect2.NewFloat32(vector2.NewFloat32(300, 200), vector2.NewFloat32(400, 10)), blocking: true, color: rl.Gray},
		{rect: rect2.NewFloat32(vector2.NewFloat32(250, 300), vector2.NewFloat32(100, 10)), blocking: true, color: rl.Gray},
		{rect: rect2.NewFloat32(vector2.NewFloat32(650, 300), vector2.NewFloat32(100, 10)), blocking: true, color: rl.Gray},
	}

	camera := rl.Camera2D{
		Offset:   vector2.Float32{X: screenWidth / 2, Y: screenHeight / 2},
		Target:   player.position,
		Rotation: 0,
		Zoom:     1,
	}

	cameraUpdaters := []cameraUpdater{
		updateCameraCenter,
		updateCameraCenterInsideMap,
		updateCameraCenterSmoothFollow,
		updateCameraEvenOutOnLanding,
		updateCameraPlayerBoundsPush,
	}
	cameraDescriptions := []string{
		"1. Follow player center",
		"2. Follow player center, but clamp to map edges",
		"3. Follow player center; smoothed",
		"4. Follow player center horizontally; update player center vertically after landing",
		"5. Player push camera on getting too close to screen edge",
	}

	cameraOption := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()

		updatePlayer(&player, envItems, deltaTime)

		camera.Zoom += rl.GetMouseWheelMove() * 0.05

		camera.Zoom = clamp(camera.Zoom, 0.25, 3)

		if rl.IsKeyPressed(rl.KeyC) {
			cameraOption = (cameraOption + 1) % len(cameraUpdaters)
		}
		// Call update camera function by its pointer
		cameraUpdaters[cameraOption](&camera, &player, envItems, deltaTime)

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(camera)

		for _, item := range envItems {
			rl.DrawRectangleRec(item.rect, item.color)
		}

		playerRect := rect2.NewFloat32(
			player.position.SubXY(20, 40),
			vector2.NewFloat32(40, 40),
		)
		rl.DrawRectangleRec(playerRect, rl.Red)
		rl.DrawCircleV(player.position, 5, rl.Gold)

		rl.EndMode2D()

		rl.DrawText("Controls:", 20, 20, 10, rl.Black)
		rl.DrawText(" - Right/Left to move", 40, 40, 10, rl.DarkGray)
		rl.DrawText(" - Space to jump", 40, 60, 10, rl.DarkGray)
		rl.DrawText(" - Mouse Wheel to Zoom in-out, R to reset zoom", 40, 80, 10, rl.DarkGray)
		rl.DrawText(" - C to change camera mode", 20, 100, 10, rl.Black)
		rl.DrawText("Current Camera Mode:", 20, 120, 10, rl.DarkGray)
		rl.DrawText(cameraDescriptions[cameraOption], 40, 140, 10, rl.DarkGray)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func updatePlayer(player *Player, envItems []EnvironmentItem, delta float32) {
	if rl.IsKeyDown(rl.KeyLeft) {
		player.position.X -= playerHORSpeed * delta
	}
	if rl.IsKeyDown(rl.KeyRight) {
		player.position.X += playerHORSpeed * delta
	}
	if rl.IsKeyDown(rl.KeySpace) && player.canJump {
		player.speed = -playerJumpSpeed
		player.canJump = false
	}

	hitObstacle := false
	for _, item := range envItems {
		if item.blocking &&
			item.rect.X() <= player.position.X && item.rect.X()+item.rect.Width() >= player.position.X &&
			item.rect.Y() >= player.position.Y && item.rect.Y() <= player.position.Y+player.speed*delta {
			hitObstacle = true
			player.speed = 0
			player.position.Y = item.rect.Y()
			break
		}
	}

	if !hitObstacle {
		player.position.Y += player.speed * delta
		player.speed += gravity * delta
		player.canJump = false
	} else {
		player.canJump = true
	}
}

func updateCameraCenter(camera *rl.Camera2D, player *Player, _ []EnvironmentItem, _ float32) {
	camera.Offset = vector2.Float32{X: screenWidth / 2, Y: screenHeight / 2}
	camera.Target = player.position
}

func updateCameraCenterInsideMap(camera *rl.Camera2D, player *Player, envItems []EnvironmentItem, _ float32) {
	camera.Offset = vector2.Float32{X: screenWidth / 2, Y: screenHeight / 2}
	camera.Target = player.position

	var minX, minY, maxX, maxY float32 = 1000, 1000, -1000, -10000

	for _, item := range envItems {
		minX = min(item.rect.X(), minX)
		maxX = max(item.rect.X()+item.rect.Width(), maxX)
		minY = min(item.rect.Y(), minY)
		maxY = max(item.rect.Y()+item.rect.Height(), maxY)
	}

	maxV := rl.GetWorldToScreen2D(vector2.Float32{X: maxX, Y: maxY}, *camera)
	minV := rl.GetWorldToScreen2D(vector2.Float32{X: minX, Y: minY}, *camera)

	if maxV.X < screenWidth {
		camera.Offset.X = screenWidth - (maxV.X - screenWidth/2)
	}
	if maxV.Y < screenHeight {
		camera.Offset.Y = screenHeight - (maxV.Y - screenHeight/2)
	}
	if minV.X > 0 {
		camera.Offset.X = screenWidth/2 - minV.X
	}
	if minV.Y > 0 {
		camera.Offset.Y = screenHeight/2 - minV.Y
	}
}

func updateCameraCenterSmoothFollow(camera *rl.Camera2D, player *Player, _ []EnvironmentItem, delta float32) {
	var minSpeed, minEffectLength, fractionSpeed float32 = 30.0, 10.0, 0.8

	camera.Offset = vector2.Float32{X: screenWidth / 2, Y: screenHeight / 2}
	diff := player.position.Sub(camera.Target)
	length := diff.LengthF()

	if length > minEffectLength {
		speed := max(fractionSpeed*length, minSpeed)
		camera.Target = camera.Target.Add(diff.ScaleF(speed * delta / length))
	}
}

func updateCameraEvenOutOnLanding(camera *rl.Camera2D, player *Player, _ []EnvironmentItem, delta float32) {
	camera.Offset = vector2.Float32{X: screenWidth / 2, Y: screenHeight / 2}
	camera.Target.X = player.position.X

	if eveningOut {
		if evenOutTarget > camera.Target.Y {
			camera.Target.Y += evenOutSpeed * delta
			if camera.Target.Y > evenOutTarget {
				camera.Target.Y = evenOutTarget
				eveningOut = false
			}
		} else {
			camera.Target.Y -= evenOutSpeed * delta
			if camera.Target.Y < evenOutTarget {
				camera.Target.Y = evenOutTarget
				eveningOut = false
			}
		}
	} else {
		if player.canJump && player.speed == 0 && player.position.Y != camera.Target.Y {
			eveningOut = true
			evenOutTarget = player.position.Y
		}
	}
}

func updateCameraPlayerBoundsPush(camera *rl.Camera2D, player *Player, _ []EnvironmentItem, _ float32) {
	bbox := vector2.Float32{X: 0.2, Y: 0.2}

	bboxWorldMin := rl.GetScreenToWorld2D(vector2.Float32{X: (1 - bbox.X) * 0.5 * screenWidth, Y: (1 - bbox.Y) * 0.5 * screenHeight}, *camera)
	bboxWorldMax := rl.GetScreenToWorld2D(vector2.Float32{X: (1 + bbox.X) * 0.5 * screenWidth, Y: (1 + bbox.Y) * 0.5 * screenHeight}, *camera)
	camera.Offset = vector2.Float32{X: (1 - bbox.X) * 0.5 * screenWidth, Y: (1 - bbox.Y) * 0.5 * screenHeight}

	if player.position.X < bboxWorldMin.X {
		camera.Target.X = player.position.X
	}
	if player.position.Y < bboxWorldMin.Y {
		camera.Target.Y = player.position.Y
	}
	if player.position.X > bboxWorldMax.X {
		camera.Target.X = bboxWorldMin.X + (player.position.X - bboxWorldMax.X)
	}
	if player.position.Y > bboxWorldMax.Y {
		camera.Target.Y = bboxWorldMin.Y + (player.position.Y - bboxWorldMax.Y)
	}
}

func clamp(zoom float32, min float32, max float32) float32 {
	if zoom < min {
		return min
	}
	if zoom > max {
		return max
	}
	return zoom
}
