package main

import "github.com/gonutz/prototype/draw"

const (
	windowW, windowH = 800, 600
)

func main() {
	x := 0
	check(draw.RunWindow("Running, out of Power", windowW, windowH, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
		}
		if window.IsKeyDown(draw.KeyLeft) {
			x--
		}
		if window.IsKeyDown(draw.KeyRight) {
			x++
		}
		window.FillRect(0, 0, windowW, windowH, draw.RGB(0.9, 0.9, 0.9))
		window.DrawImageFile("old_guy.png", x, 200)
	}))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
