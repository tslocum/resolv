package main

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenWidth int32 = 320
var screenHeight int32 = 240
var cell int32 = 4

var drawHelpText = true

func main() {

	// defer profile.Start(profile.ProfilePath(".")).Stop()

	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(screenWidth, screenHeight, "resolv Tests")

	worldIndex := 0
	worlds := []WorldInterface{
		&WorldBounce{},
		&WorldZones{},
		&WorldLines{},
		&WorldPlatformer{},
		&WorldCompound{},
		&WorldShooter{},
	}

	for _, world := range worlds {
		world.Create()
	}

	targetFPS := int32(60)
	rl.SetTargetFPS(targetFPS)

	framebuffer := rl.LoadRenderTexture(screenWidth, screenHeight) // Make a framebuffer so it stretches to fill the window
	rl.SetTextureFilter(framebuffer.Texture, rl.FilterPoint)       // No blurry!

	for !rl.WindowShouldClose() {

		world := worlds[worldIndex]

		if rl.IsKeyPressed(rl.KeyOne) {
			rl.SetWindowSize(int(screenWidth), int(screenHeight))
		} else if rl.IsKeyPressed(rl.KeyTwo) {
			rl.SetWindowSize(int(screenWidth*2), int(screenHeight*2))
		} else if rl.IsKeyPressed(rl.KeyThree) {
			rl.SetWindowSize(int(screenWidth*3), int(screenHeight*3))
		} else if rl.IsKeyPressed(rl.KeyFour) {
			rl.SetWindowSize(int(screenWidth*4), int(screenHeight*4))
		} else if rl.IsKeyPressed(rl.KeyFive) {
			rl.SetWindowSize(int(screenWidth*5), int(screenHeight*5))
		}

		if rl.IsKeyPressed(rl.KeyF2) {
			targetFPS = 60
			rl.SetTargetFPS(targetFPS)
		} else if rl.IsKeyPressed(rl.KeyF3) {
			targetFPS = 30
			rl.SetTargetFPS(targetFPS)
		} else if rl.IsKeyPressed(rl.KeyF4) {
			targetFPS = 10
			rl.SetTargetFPS(targetFPS)
		}

		if rl.IsKeyPressed(rl.KeyE) {
			worldIndex++
		} else if rl.IsKeyPressed(rl.KeyQ) {
			worldIndex--
		}

		if worldIndex > len(worlds)-1 {
			worldIndex = 0
		} else if worldIndex < 0 {
			worldIndex = len(worlds) - 1
		}

		if rl.IsKeyPressed(rl.KeyF1) {
			drawHelpText = !drawHelpText
		}

		if rl.IsKeyPressed(rl.KeyR) {
			world.Destroy()
			world.Create()
		}

		world.Update()

		rl.ClearBackground(rl.Color{20, 20, 40, 255})

		rl.BeginTextureMode(framebuffer)

		rl.BeginDrawing()

		world.Draw()

		fps := strconv.Itoa(int(math.Round(float64(1.0 / 60.0 / rl.GetFrameTime() * 60.0))))

		DrawText(screenWidth-32, 0, fps)

		if drawHelpText {
			DrawText(16, screenHeight-64,
				"Use F2, F3, and F4 to change the target framerate.",
				"Use the number keys to change the window scale.",
				"Use the 'Q' and 'E' keys to",
				"jump between worlds.",
				"Press F1 to turn on or off this text.",
			)
		}

		rl.EndTextureMode()

		rl.DrawTexturePro(framebuffer.Texture, rl.Rectangle{0, 0, float32(screenWidth), -float32(screenHeight)},
			rl.Rectangle{0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())}, rl.Vector2{}, 0, rl.White)

		rl.EndDrawing()

	}

}
