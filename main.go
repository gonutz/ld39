// In this game you have to win a running race between a lot of old people.
// You have to alternatingly press left and right to speed up.

package main

import (
	"math/rand"
	"time"

	"github.com/gonutz/prototype/draw"
)

const (
	windowW, windowH = 1600, 600
	acceleration     = 0.3
	decelration      = 0.05
	maxSpeed         = 2.5
	walkFrameDelay   = 5.0
	blinkOverlay     = "blink.png"
	shutMouthOverlay = "shut_mouth.png"
)

var (
	walkFrames = []string{
		"old_guy1.png",
		"old_guy2.png",
		"old_guy1.png",
		"old_guy3.png",
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var x float64
	var speed float64
	var nextUpLeft bool
	walkFrame := 0
	nextFrame := walkFrameDelay
	var blinkTimer int
	var mouthShutTimer int

	check(draw.RunWindow("Running, out of Power", windowW, windowH, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
		}
		left := window.IsKeyDown(draw.KeyLeft)
		right := window.IsKeyDown(draw.KeyRight)
		if nextUpLeft && left && !right ||
			!nextUpLeft && right && !left {
			speed += acceleration
			nextUpLeft = !nextUpLeft
		}
		if speed > maxSpeed {
			speed = maxSpeed
		}
		speed -= decelration
		if speed < 0 {
			speed = 0
		}
		x += speed
		nextFrame -= speed
		if nextFrame <= 0 {
			nextFrame = walkFrameDelay
			walkFrame = (walkFrame + 1) % len(walkFrames)
		}
		if speed == 0 {
			walkFrame = 0
		}
		blinkTimer--
		if blinkTimer < -3 {
			blinkTimer = 30 + rand.Intn(90)
		}
		window.FillRect(0, 0, windowW, windowH, draw.RGB(0.9, 0.9, 0.9))
		window.DrawImageFile(walkFrames[walkFrame], int(x+0.5), 200)
		mouthShutTimer--
		if mouthShutTimer < 0 {
			mouthShutTimer = 0
		}
		if speed < 1 {
			if mouthShutTimer == 0 {
				window.DrawImageFile(shutMouthOverlay, int(x+0.5), 200)
			}
		} else {
			mouthShutTimer = 10
		}
		if blinkTimer <= 0 {
			window.DrawImageFile(blinkOverlay, int(x+0.5), 200)
		}
	}))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
