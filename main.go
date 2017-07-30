// In this game you have to win a running race between a lot of old people.
// You have to alternatingly press left and right to speed up.

package main

import (
	"math/rand"
	"time"

	"github.com/gonutz/prototype/draw"
)

const (
	windowW, windowH       = 1200, 600
	acceleration           = maxSpeed //0.3
	decelration            = 0.05
	maxSpeed               = 2.5
	walkFrameDelay         = 5.0
	blinkOverlay           = "blink.png"
	shutMouthOverlay       = "shut_mouth.png"
	doorShut               = "closed_door.png"
	nurseX, nurseY         = 600, windowH - backTilesH - 309
	tv                     = "tv.png"
	tvX, tvY               = 1050, 230
	backTiles              = "back_tiles.png"
	backTilesW, backTilesH = 143, 218
	table                  = "table.png"
	tableX, tableY         = 870, 375
	goalX                  = 735.0
	armchair               = "armchair.png"
	armchairX, armchairY   = 0, 315
	couch                  = "couch.png"
	couchX, couchY         = 160, 245
	couchBack              = "couch_back.png"
	couchBackX, couchBackY = 140, 395
)

var (
	walkFrames = []string{
		"old_guy1.png",
		"old_guy2.png",
		"old_guy1.png",
		"old_guy3.png",
	}
	nurseFrames = []string{
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse1.png",
		"nurse3.png",
		"nurse2.png",
		"nurse1.png",
		"nurse2.png",
		"nurse1.png",
		"nurse3.png",
		"nurse2.png",
		"nurse3.png",
		"nurse1.png",
		"nurse2.png",
		"nurse1.png",
		"nurse3.png",
		"nurse2.png",
		"nurse3.png",
		"nurse1.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	x := 130.0
	var speed float64
	var nextUpLeft bool
	walkFrame := 0
	nextFrame := walkFrameDelay
	var blinkTimer int
	var mouthShutTimer int
	nurse := nurseFrames
	var nurseTalking bool
	nurseTimer := 0
	nurseAnimationStart := 120
	cameraX := 0

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
		if x > goalX {
			x = goalX
			speed = 0
		}
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
		mouthShutTimer--
		if mouthShutTimer < 0 {
			mouthShutTimer = 0
		}
		nurseAnimationStart--
		if nurseAnimationStart < 0 {
			nurseAnimationStart = 0
		}
		if nurseAnimationStart == 1 {
			nurseTalking = true
		}
		cameraX = int(x+0.5) - windowW/4
		if cameraX < 0 {
			cameraX = 0
		}

		// render scene

		// clear background
		window.FillRect(0, 0, windowW, windowH, draw.RGB(0.9, 0.9, 0.9))
		// draw floor
		for i := 0; i < 20; i++ {
			window.DrawImageFile(backTiles, i*backTilesW-cameraX, windowH-backTilesH)
		}
		// draw nurse
		if nurseTalking {
			window.DrawImageFile(nurse[0], nurseX-cameraX, nurseY)
			nurseTimer--
			if nurseTimer <= 0 {
				nurseTimer = 10
				nurse = nurse[1:]
				if len(nurse) == 0 {
					nurseTalking = false
				}
			}
		} else {
			window.DrawImageFile(doorShut, nurseX-cameraX, nurseY)
		}
		// draw armchair and couch in the background
		window.DrawImageFile(couch, couchX-cameraX, couchY)
		window.DrawImageFile(armchair, armchairX-cameraX, armchairY)
		// draw main guy
		window.DrawImageFile(walkFrames[walkFrame], int(x+0.5)-cameraX, 200)
		if speed < 1 {
			if mouthShutTimer == 0 {
				window.DrawImageFile(shutMouthOverlay, int(x+0.5)-cameraX, 200)
			}
		} else {
			mouthShutTimer = 10
		}
		if blinkTimer <= 0 {
			window.DrawImageFile(blinkOverlay, int(x+0.5)-cameraX, 200)
		}
		// draw couch in the foreground
		window.DrawImageFile(couchBack, couchBackX-cameraX, couchBackY)
		// draw TV set
		window.DrawImageFile(table, tableX-cameraX, tableY)
		window.DrawImageFile(tv, tvX-cameraX, tvY)
	}))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
