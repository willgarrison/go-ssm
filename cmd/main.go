package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/willgarrison/go-ssm/pkg/ui"
)

var (
	windowRect pixel.Rect = pixel.R(0, 0, 940, 940)
	ssmRect    pixel.Rect = pixel.R(60.1, 60.1, 880, 880)
)

func main() {
	pixelgl.Run(run)
}

func run() {

	// midiOutput, err := output.NewMidiOutput()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer midiOutput.Driver.Close()

	pattern := []string{
		"E", "G", "B", "B", "B",
		"A", "C", "B", "A", "G", "G",
		"A", "B", "G", "E", "G", "A", "F#", "E", "D", "E", "E",
	}

	// pattern := []string{
	// 	"C", "C", "F", "F", "C", "C", "G", "G",
	// 	"C", "C", "C", "C", "C", "C", "G", "C",
	// }

	// UI
	win := ui.NewWindow(windowRect)
	ssm := ui.NewSSM(ssmRect, pattern)

	imdBatch := imdraw.New(nil)

	for !win.Closed() {

		// Clear
		win.Clear(ui.ColorBackground)
		imdBatch.Clear()

		// Update
		ssm.Update(win)

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
