package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/willgarrison/go-iobox"
)

type InputGrid struct {
	IOBox       *iobox.IOBox
	Signal      iobox.Signal
	Imd         *imdraw.IMDraw
	Rect        pixel.Rect
	Grid        [][]uint32
	lastMouseX  int32
	lastMouseY  int32
	isComposing bool
}

func NewInputGrid(rect pixel.Rect) *InputGrid {

	inputGrid := &InputGrid{
		IOBox:      iobox.New(),
		Imd:        imdraw.New(nil),
		Rect:       rect,
		lastMouseX: -1,
		lastMouseY: -1,
	}

	inputGrid.Compose()

	return inputGrid
}

func (inputGrid *InputGrid) Compose() {

	xSteps := 16
	ySteps := 12

	inputGrid.isComposing = true

	inputGrid.Imd.Clear()

	inputGrid.Grid = make([][]uint32, xSteps)
	for i := range inputGrid.Grid {
		inputGrid.Grid[i] = make([]uint32, 12)
	}

	blockWidth := inputGrid.Rect.W() / float64(xSteps)
	blockHeight := inputGrid.Rect.H() / float64(ySteps)

	// Vertical Lines
	for x := 0; x <= xSteps; x++ {
		inputGrid.Imd.Color = ColorGridLines
		inputGrid.Imd.Push(
			pixel.V(
				inputGrid.Rect.Min.X+float64(x)*blockWidth,
				inputGrid.Rect.Min.Y,
			),
			pixel.V(
				inputGrid.Rect.Min.X+float64(x)*blockWidth,
				inputGrid.Rect.Max.Y,
			),
		)
		inputGrid.Imd.Line(1)
	}

	// Horizontal Lines
	for y := 0; y <= ySteps; y++ {
		inputGrid.Imd.Color = ColorGridLines
		inputGrid.Imd.Push(
			pixel.V(
				inputGrid.Rect.Min.X,
				inputGrid.Rect.Min.Y+(float64(y)*blockHeight),
			),
			pixel.V(
				inputGrid.Rect.Max.X,
				inputGrid.Rect.Min.Y+(float64(y)*blockHeight),
			),
		)
		inputGrid.Imd.Line(1)
	}

	inputGrid.isComposing = false
}

func (inputGrid *InputGrid) Update(win *pixelgl.Window) {

}

func (inputGrid *InputGrid) DrawTo(imd *imdraw.IMDraw) {
	if !inputGrid.isComposing {
		inputGrid.Imd.Draw(imd)
	}
}
