// In this game you have to win a running race between a lot of old people.
// You have to alternatingly press left and right to speed up.

package main

import (
	"math/rand"
	"time"

	"github.com/gonutz/prototype/draw"
)

type gameState int

const (
	outsideFadingIn gameState = iota
	outsideFadingOut
	insideFadingIn
	waitingForNurse
	nurseTalks
	runningRace
)

const (
	windowW, windowH       = 1200, 600
	acceleration           = 0.25 //maxSpeed //0.25
	decelration            = 0.05
	maxSpeed               = 2.5
	walkFrameDelay         = 5.0
	blinkOverlay           = "blink.png"
	shutMouthOverlay       = "shut_mouth.png"
	doorShut               = "closed_door.png"
	nurseX, nurseY         = 900, windowH - backTilesH - 309
	tv                     = "tv.png"
	tvX, tvY               = 1550, 230
	backTiles              = "back_tiles.png"
	backTilesW, backTilesH = 143, 218
	table                  = "table.png"
	tableX, tableY         = 1370, 375
	goalX                  = 1235.0
	armchair               = "armchair.png"
	armchairX, armchairY   = 0, 315
	couch                  = "couch.png"
	couchX, couchY         = 200, 245
	couchBack              = "couch_back.png"
	couchBackX, couchBackY = 190, 395
	sitting                = "old_guy_sitting.png"
	sittingBlinkOverlay    = "old_guy_sitting_blink.png"
	sittingX, sittingY     = -3, 195
	sceneW                 = 1800
	painting               = "painting1.png"
	paintingX, paintingY   = 700, 70
	painting2              = "painting2.png"
	painting2X, painting2Y = 1500, 50
	background             = "outside.png"
	startBlend             = 1.1
	dBlendSlow             = -0.005
	dBlendFast             = 3 * dBlendSlow
	endBlend               = -1.3
	bird1                  = "bird1.wav"
	bird2                  = "bird2.wav"
	bird3                  = "bird3.wav"
	bird4                  = "bird4.wav"
	birdLeftUp             = "bird_left_up.png"
	birdLeftDown           = "bird_left_down.png"
	birdRightUp            = "bird_right_up.png"
	birdRightDown          = "bird_right_down.png"
	nurseSpeech            = "nurse.wav"
	woman                  = "old_broad.png"
	man                    = "other_dude.png"
	squeak                 = "squeak.wav"
	dSqueak                = 74
	hitTable               = "hit_table.wav"
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
		"nurse1.png",
		"nurse1.png",
		"nurse3.png",
		"nurse1.png",
		"nurse2.png",
		"nurse1.png",
		"nurse3.png",
		"nurse2.png",
		"nurse3.png",
		"nurse1.png",
		"nurse2.png",
		"nurse1.png",
		"nurse2.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
		"nurse3.png",
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var state gameState = runningRace //waitingForNurse
	x := 130.0
	nextSqueak := int(x + 1)
	var goalReached bool
	var speed float64
	var nextUpLeft bool
	walkFrame := 0
	nextFrame := walkFrameDelay
	var blinkTimer int
	var mouthShutTimer int
	nurse := nurseFrames
	//var nurseTalking bool
	nurseTimer := 0
	nurseAnimationStart := 2
	cameraX := 0
	var blend float32 = startBlend
	var stateTimer int
	type bird struct {
		x, y     int
		left, up bool
	}
	birds := []bird{
		bird{x: 200, y: 450, left: false, up: true},
		bird{x: 80, y: 200, left: false, up: false},
		bird{x: 680, y: 450, left: true, up: true},
		bird{x: 850, y: 350, left: true, up: true},
	}
	birdToggleTimer := 0

	check(draw.RunWindow("Running, out of Power", windowW, windowH, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
		}

		if state == runningRace {
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
			if x >= float64(nextSqueak) {
				window.PlaySoundFile(squeak)
				nextSqueak += dSqueak
			}
			if x > goalX {
				x = goalX
				speed = 0
				if !goalReached {
					window.PlaySoundFile(hitTable)
				}
				goalReached = true
			}
			nextFrame -= speed
			if nextFrame <= 0 {
				nextFrame = walkFrameDelay
				walkFrame = (walkFrame + 1) % len(walkFrames)
			}
			if speed == 0 {
				walkFrame = 0
			}
			mouthShutTimer--
			if mouthShutTimer < 0 {
				mouthShutTimer = 0
			}
		}
		blinkTimer--
		if blinkTimer < -3 {
			blinkTimer = 30 + rand.Intn(90)
		}
		nurseAnimationStart--
		if nurseAnimationStart < 0 {
			nurseAnimationStart = 0
		}
		//if nurseAnimationStart == 1 {
		//	nurseTalking = true
		//}
		// fix camera on character
		cameraX = int(x+0.5) - windowW/4
		if cameraX < 0 {
			cameraX = 0
		}
		maxCamX := sceneW - windowW
		if cameraX > maxCamX {
			cameraX = maxCamX
		}
		if state == waitingForNurse && stateTimer > 150 {
			state = nurseTalks
			stateTimer = 0
			window.PlaySoundFile(nurseSpeech)
		}
		birdToggleTimer--
		if birdToggleTimer < 0 {
			birdToggleTimer = rand.Intn(30)
			i := rand.Intn(len(birds))
			b := &birds[i]
			b.up = !b.up
		}
		stateTimer++

		// render scene

		renderOutside := func() {
			window.DrawImageFile(background, 0, 0)
			// add birds
			for _, bird := range birds {
				var img string
				if bird.left {
					if bird.up {
						img = birdLeftUp
					} else {
						img = birdLeftDown
					}
				} else {
					if bird.up {
						img = birdRightUp
					} else {
						img = birdRightDown
					}
				}
				window.DrawImageFile(img, bird.x, bird.y)
			}
		}

		renderInside := func() {
			// clear background
			window.FillRect(0, 0, windowW, windowH, draw.RGB(0.9, 0.9, 0.9))
			// draw nice things on the wall
			window.DrawImageFile(painting, paintingX-cameraX, paintingY)
			window.DrawImageFile(painting2, painting2X-cameraX, painting2Y)
			// draw floor
			for i := 0; i < 20; i++ {
				window.DrawImageFile(backTiles, i*backTilesW-cameraX, windowH-backTilesH)
			}
			// TODO remove:
			window.DrawImageFile(woman, 600-cameraX, 150)
			window.DrawImageFile(man, 700-cameraX, 250)
		}

		if state == outsideFadingIn {
			// make the birds peep
			switch stateTimer {
			case 310:
				window.PlaySoundFile(bird1)
			case 100, 200:
				window.PlaySoundFile(bird2)
			case 290, 440:
				window.PlaySoundFile(bird3)
			case 150, 360:
				window.PlaySoundFile(bird4)
			}
			// fade in
			renderOutside()
			a := blend
			if a < 0 {
				a = 0
			}
			if a > 1 {
				a = 1
			}
			window.FillRect(0, 0, windowW, windowH, draw.RGBA(0, 0, 0, a))
			blend += dBlendSlow
			if blend < endBlend {
				state = outsideFadingOut
				blend = 0
			}
		} else if state == outsideFadingOut {
			renderOutside()
			window.FillRect(0, 0, windowW, windowH, draw.RGBA(0, 0, 0, blend))
			blend -= dBlendFast
			if blend > 1 {
				state = insideFadingIn
				blend = 1
			}
		} else if state == insideFadingIn {
			renderInside()
			window.DrawImageFile(doorShut, nurseX-cameraX, nurseY)
			// draw armchair and couches
			window.DrawImageFile(couch, couchX-cameraX, couchY)
			window.DrawImageFile(sitting, sittingX-cameraX, sittingY)
			if blinkTimer <= 0 {
				window.DrawImageFile(sittingBlinkOverlay, sittingX-cameraX, sittingY)
			}
			window.DrawImageFile(couchBack, couchBackX-cameraX, couchBackY)
			// draw TV set
			window.DrawImageFile(table, tableX-cameraX, tableY)
			window.DrawImageFile(tv, tvX-cameraX, tvY)
			// fade in
			window.FillRect(0, 0, windowW, windowH, draw.RGBA(0, 0, 0, blend))
			blend += dBlendFast
			if blend < 0 {
				state = waitingForNurse
				stateTimer = 0
				blend = 0
			}
		} else if state == waitingForNurse {
			renderInside()
			window.DrawImageFile(doorShut, nurseX-cameraX, nurseY)
			// draw armchair and couches
			window.DrawImageFile(couch, couchX-cameraX, couchY)
			window.DrawImageFile(sitting, sittingX-cameraX, sittingY)
			if blinkTimer <= 0 {
				window.DrawImageFile(sittingBlinkOverlay, sittingX-cameraX, sittingY)
			}
			window.DrawImageFile(couchBack, couchBackX-cameraX, couchBackY)
			// draw TV set
			window.DrawImageFile(table, tableX-cameraX, tableY)
			window.DrawImageFile(tv, tvX-cameraX, tvY)
		} else if state == nurseTalks {
			renderInside()
			// draw nurse
			window.DrawImageFile(nurse[0], nurseX-cameraX, nurseY)
			nurseTimer--
			if nurseTimer <= 0 {
				nurseTimer = 10
				nurse = nurse[1:]
				if len(nurse) == 0 {
					state = runningRace // TODO
				}
			}
			// draw armchair and couch in the background
			window.DrawImageFile(couch, couchX-cameraX, couchY)
			window.DrawImageFile(armchair, armchairX-cameraX, armchairY)
			// draw main guy
			window.DrawImageFile(sitting, sittingX-cameraX, sittingY)
			if blinkTimer <= 0 {
				window.DrawImageFile(sittingBlinkOverlay, sittingX-cameraX, sittingY)
			}
			// draw couch in the foreground
			window.DrawImageFile(couchBack, couchBackX-cameraX, couchBackY)
			// draw TV set
			window.DrawImageFile(table, tableX-cameraX, tableY)
			window.DrawImageFile(tv, tvX-cameraX, tvY)
		} else if state == runningRace {
			renderInside()
			window.DrawImageFile(doorShut, nurseX-cameraX, nurseY)
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
		}
	}))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
