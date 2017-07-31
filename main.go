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
	getReadyForRace
	runningRace
)

const (
	windowW, windowH           = 1200, 600
	acceleration               = 0.25
	decelration                = 0.05
	maxSpeed                   = 2.1
	walkFrameDelay             = 5.0
	blinkOverlay               = "blink.png"
	shutMouthOverlay           = "shut_mouth.png"
	doorShut                   = "closed_door.png"
	nurseX, nurseY             = 900, windowH - backTilesH - 309
	tv                         = "tv.png"
	tvX, tvY                   = 1550, 230
	backTiles                  = "back_tiles.png"
	backTilesW, backTilesH     = 143, 218
	table                      = "table.png"
	tableEmpty                 = "table_empty.png"
	tableX, tableY             = 1370, 375
	goalX                      = 1235.0
	armchair                   = "armchair.png"
	armchairX, armchairY       = 0, 315
	couch                      = "couch.png"
	couchX, couchY             = 200, 245
	sleepyWoman                = "sleeping_woman.png"
	sleepyWomanX, sleepyWomanY = 200, 245 - 110
	couchBack                  = "couch_back.png"
	couchBackX, couchBackY     = 190, 395
	manSitting                 = "other_dude_sitting.png"
	manSittingX, manSittingY   = 190, 395 - 60
	sitting                    = "old_guy_sitting.png"
	sittingBlinkOverlay        = "old_guy_sitting_blink.png"
	sittingX, sittingY         = -3, 195
	sceneW                     = 1800
	painting                   = "painting1.png"
	paintingX, paintingY       = 700, 70
	painting2                  = "painting2.png"
	painting2X, painting2Y     = 1500, 50
	background                 = "outside.png"
	startBlend                 = 1.1
	dBlendSlow                 = -0.005
	dBlendFast                 = 3 * dBlendSlow
	endBlend                   = -1.3
	bird1                      = "bird1.wav"
	bird2                      = "bird2.wav"
	bird3                      = "bird3.wav"
	bird4                      = "bird4.wav"
	birdLeftUp                 = "bird_left_up.png"
	birdLeftDown               = "bird_left_down.png"
	birdRightUp                = "bird_right_up.png"
	birdRightDown              = "bird_right_down.png"
	nurseSpeech                = "nurse.wav"
	woman                      = "old_broad.png"
	man                        = "other_dude.png"
	squeak                     = "squeak.wav"
	dSqueak                    = 74
	hitTable                   = "hit_table.wav"
	arrowLeft, arrowRight      = "instruction_left.png", "instruction_right.png"
	arrowX, arrowY             = windowW/2 - 72/2, 0
	arrowToggleTime            = 20
	manAcceleration            = 0.003
	maxManSpeed                = 2
	manGoalX                   = 1255.0
	mansChoice                 = "matlock.wav"
	manWins                    = "other_guy_wins.wav"
	womanAcceleration          = 0.004
	maxWomanSpeed              = 1.55
	womanGoalX                 = 1335.0
	womansChoice               = "grays_anatomy.wav"
	backMusic                  = "back_music.wav"
)

var (
	walkFrames = []string{
		"old_guy1.png",
		"old_guy2.png",
		"old_guy1.png",
		"old_guy3.png",
	}
	manWalkFrames = []string{
		"other_dude1.png",
		"other_dude2.png",
		"other_dude1.png",
		"other_dude3.png",
	}
	manWinning      = "other_dude_winning.png"
	manTalkOverlay  = "other_dude_mouth_talks.png"
	womanWalkFrames = []string{
		"woman1.png",
		"woman2.png",
		"woman1.png",
		"woman3.png",
	}
	womanTalkOverlay = "old_broad_talk_mouth.png"
	nurseFrames      = []string{
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

	var state gameState = getReadyForRace //= waitingForNurse
	x := 130.0
	manX, manY := 550.0, 250
	manSpeed := 0.0
	var playerWon, manWon bool
	womanX, womanY := 530.0, 150
	womanSpeed := 0.0
	nextSqueak := int(x + 1)
	var speed float64
	var nextUpLeft bool
	walkFrame := 0
	nextFrame := walkFrameDelay
	manWalkFrame := 0
	manNextFrame := walkFrameDelay
	womanWalkFrame := 0
	womanNextFrame := walkFrameDelay
	var blinkTimer int
	var mouthShutTimer int
	nurse := nurseFrames
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
	arrowTimer := 0
	var showArrowLeft bool
	var playerKnowsHowToPlay bool
	musicTimer := 0
	var musicStarted bool

	resetRace := func() {
		state = getReadyForRace
		x = 130.0
		manX, manY = 550.0, 250
		manSpeed = 0.0
		playerWon, manWon = false, false
		womanX, womanY = 530.0, 150
		womanSpeed = 0.0
		nextSqueak = int(x + 1)
		speed = 0
		nextUpLeft = false
		walkFrame = 0
		nextFrame = walkFrameDelay
		manWalkFrame = 0
		manNextFrame = walkFrameDelay
		womanWalkFrame = 0
		womanNextFrame = walkFrameDelay
		mouthShutTimer = 0
		cameraX = 0
		stateTimer = 0
		showArrowLeft = false
	}
	resetRace()

	check(draw.RunWindow("Running, out of Power", windowW, windowH, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
		}

		if musicStarted {
			if musicTimer == 0 {
				window.PlaySoundFile(backMusic)
				musicTimer = 100
			}
			musicTimer--
		}

		if state == runningRace {
			musicStarted = true
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
				if !playerWon {
					window.PlaySoundFile(hitTable)
				}
				playerWon = true
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
			// move man
			manSpeed += manAcceleration
			if manSpeed > maxManSpeed {
				manSpeed = maxManSpeed
			}
			manX += manSpeed
			if manX > manGoalX {
				manX = manGoalX
				manSpeed = 0
				if !manWon && !playerWon {
					window.PlaySoundFile(manWins)
					manWon = true
				}
			}
			manNextFrame -= manSpeed
			if manNextFrame <= 0 {
				manNextFrame = walkFrameDelay
				manWalkFrame = (manWalkFrame + 1) % len(manWalkFrames)
			}
			// move woman
			womanSpeed += womanAcceleration
			if womanSpeed > maxWomanSpeed {
				womanSpeed = maxWomanSpeed
			}
			womanX += womanSpeed
			if womanX > womanGoalX {
				womanX = womanGoalX
				womanSpeed = 0
			}
			womanNextFrame -= womanSpeed
			if womanNextFrame <= 0 {
				womanNextFrame = walkFrameDelay
				womanWalkFrame = (womanWalkFrame + 1) % len(womanWalkFrames)
			}

			if x > 500 {
				playerKnowsHowToPlay = true
			}

			if window.WasKeyPressed(draw.KeySpace) {
				resetRace()
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
			window.DrawImageFile(sleepyWoman, sleepyWomanX-cameraX, sleepyWomanY)
			window.DrawImageFile(sitting, sittingX-cameraX, sittingY)
			if blinkTimer <= 0 {
				window.DrawImageFile(sittingBlinkOverlay, sittingX-cameraX, sittingY)
			}
			window.DrawImageFile(manSitting, manSittingX-cameraX, manSittingY)
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
			window.DrawImageFile(sleepyWoman, sleepyWomanX-cameraX, sleepyWomanY)
			window.DrawImageFile(sitting, sittingX-cameraX, sittingY)
			if blinkTimer <= 0 {
				window.DrawImageFile(sittingBlinkOverlay, sittingX-cameraX, sittingY)
			}
			window.DrawImageFile(manSitting, manSittingX-cameraX, manSittingY)
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
					state = getReadyForRace
					stateTimer = 0
				}
			}
			// draw armchair and couch in the background
			window.DrawImageFile(sleepyWoman, sleepyWomanX-cameraX, sleepyWomanY)
			window.DrawImageFile(armchair, armchairX-cameraX, armchairY)
			// draw main guy
			window.DrawImageFile(sitting, sittingX-cameraX, sittingY)
			if blinkTimer <= 0 {
				window.DrawImageFile(sittingBlinkOverlay, sittingX-cameraX, sittingY)
			}
			// draw couch in the foreground
			window.DrawImageFile(manSitting, manSittingX-cameraX, manSittingY)
			// draw TV set
			window.DrawImageFile(table, tableX-cameraX, tableY)
			window.DrawImageFile(tv, tvX-cameraX, tvY)
		} else if state == getReadyForRace {
			renderInside()
			window.DrawImageFile(doorShut, nurseX-cameraX, nurseY)
			// draw armchair and couch in the background
			window.DrawImageFile(couch, couchX-cameraX, couchY)
			window.DrawImageFile(armchair, armchairX-cameraX, armchairY)
			// draw woman in the background
			window.DrawImageFile(womanWalkFrames[0], int(womanX+0.5)-cameraX, womanY)
			if 90 <= stateTimer && stateTimer <= 165 && (stateTimer/10)%2 == 1 {
				window.DrawImageFile(womanTalkOverlay, int(womanX+0.5)-cameraX, womanY)
			}
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
			// draw guy in the front
			window.DrawImageFile(manWalkFrames[0], int(manX+0.5)-cameraX, manY)
			if 30 <= stateTimer && stateTimer <= 90 && (stateTimer/10)%2 == 0 {
				window.DrawImageFile(manTalkOverlay, int(manX+0.5)-cameraX, manY)
			}
			// draw couch in the foreground
			window.DrawImageFile(couchBack, couchBackX-cameraX, couchBackY)
			// draw TV set
			if manWon || playerWon {
				window.DrawImageFile(tableEmpty, tableX-cameraX, tableY)
			} else {
				window.DrawImageFile(table, tableX-cameraX, tableY)
			}
			window.DrawImageFile(tv, tvX-cameraX, tvY)
			if stateTimer == 30 {
				window.PlaySoundFile(mansChoice)
			}
			if stateTimer == 90 {
				window.PlaySoundFile(womansChoice)
			}
			if stateTimer > 170 {
				state = runningRace
			}
		} else if state == runningRace {
			renderInside()
			if !playerKnowsHowToPlay {
				if showArrowLeft {
					window.DrawImageFile(arrowLeft, arrowX, arrowY)
				} else {
					window.DrawImageFile(arrowRight, arrowX, arrowY)
				}
				arrowTimer--
				if arrowTimer < 0 {
					arrowTimer = arrowToggleTime
					showArrowLeft = !showArrowLeft
				}
			}
			window.DrawImageFile(doorShut, nurseX-cameraX, nurseY)
			// draw armchair and couch in the background
			window.DrawImageFile(couch, couchX-cameraX, couchY)
			window.DrawImageFile(armchair, armchairX-cameraX, armchairY)
			// draw woman in the background
			window.DrawImageFile(womanWalkFrames[womanWalkFrame], int(womanX+0.5)-cameraX, womanY)
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
			// draw guy in the front
			if manWon {
				window.DrawImageFile(manWinning, int(manX+0.5)-cameraX, manY)
			} else {
				window.DrawImageFile(manWalkFrames[manWalkFrame], int(manX+0.5)-cameraX, manY)
			}
			// draw couch in the foreground
			window.DrawImageFile(couchBack, couchBackX-cameraX, couchBackY)
			// draw TV set
			if manWon {
				window.DrawImageFile(tableEmpty, tableX-cameraX, tableY)
			} else {
				window.DrawImageFile(table, tableX-cameraX, tableY)
			}
			window.DrawImageFile(tv, tvX-cameraX, tvY)
		}
	}))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
