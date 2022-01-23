package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/willgarrison/go-ssm/pkg/ui"
)

var (
	windowRect pixel.Rect = pixel.R(0, 0, 1200, 900)
	ssmRect    pixel.Rect = pixel.R(60.1, 60.1, 1140, 840)
)

func main() {
	pixelgl.Run(run)
}

func run() {

	pattern := []rune{}

	// UI
	win := ui.NewWindow("Self-similarity Matrix (SSM)", windowRect)
	ssm := ui.NewSSM(ssmRect, pattern)
	imdBatch := imdraw.New(nil)

	for !win.Closed() {

		// Clear
		win.Clear(ui.ColorBackground)
		imdBatch.Clear()

		// Get typed input
		for _, r := range win.Typed() {
			pattern = append(pattern, r)
		}
		// Listen for backspace
		if win.JustPressed(pixelgl.KeyBackspace) || win.Repeated(pixelgl.KeyBackspace) {
			if len(pattern) > 0 {
				pattern = pattern[:len(pattern)-1]
			}
		}
		// Listen for escape
		if win.JustPressed(pixelgl.KeyEscape) {
			pattern = []rune{}
		}

		// If the new pattern is not equal to the old pattern, update the SSM
		if patternsAreNotEqual(ssm.Pattern, pattern) {
			ssm.Update(pattern)
		}

		// Draw to batch
		ssm.DrawTo(imdBatch)

		// Draw to window buffer
		imdBatch.Draw(win)

		// Draw text to window buffer
		ssm.DrawTextTo(win)

		// Update window
		win.Update()
	}
}

func patternsAreNotEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return true
	}
	for x := range a {
		if a[x] != b[x] {
			return true
		}
	}
	return false
}
