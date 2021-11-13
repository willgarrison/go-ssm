package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/willgarrison/go-ssm/pkg/output"
	"github.com/willgarrison/go-ssm/pkg/ui"
)

var (
	windowRect     pixel.Rect = pixel.R(0, 0, 1200, 900)
	inputGridRect  pixel.Rect = pixel.R(200.1, 0.1, 1200, 300)
	outputGridRect pixel.Rect = pixel.R(0.1, 300.1, 1200, 900)
)

func main() {
	pixelgl.Run(run)
}

func run() {

	midiOutput, err := output.NewMidiOutput()
	if err != nil {
		panic(err.Error())
	}
	defer midiOutput.Driver.Close()

	// UI
	win := ui.NewWindow(windowRect)
	inputGrid := ui.NewInputGrid(inputGridRect)
	outputGrid := ui.NewOutputGrid(outputGridRect)

	imdBatch := imdraw.New(nil)

	for !win.Closed() {

		// Clear
		win.Clear(ui.ColorBackground)
		imdBatch.Clear()

		// Update
		inputGrid.Update(win)
		outputGrid.Update(win)

		// Draw to batch
		inputGrid.DrawTo(imdBatch)
		outputGrid.DrawTo(imdBatch)

		// Draw to window buffer
		imdBatch.Draw(win)

		// Update window
		win.Update()
	}
}
