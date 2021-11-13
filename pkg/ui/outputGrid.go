package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/willgarrison/go-iobox"
)

type OutputGrid struct {
	IOBox       *iobox.IOBox
	Signal      iobox.Signal
	Imd         *imdraw.IMDraw
	Rect        pixel.Rect
	Grid        [][]uint32
	lastMouseX  int32
	lastMouseY  int32
	isComposing bool
}

func NewOutputGrid(rect pixel.Rect) *OutputGrid {

	outputGrid := &OutputGrid{
		IOBox:      iobox.New(),
		Imd:        imdraw.New(nil),
		Rect:       rect,
		lastMouseX: -1,
		lastMouseY: -1,
	}

	outputGrid.Compose()

	return outputGrid
}

func (outputGrid *OutputGrid) Compose() {

	xSteps := 16
	ySteps := 12

	outputGrid.isComposing = true

	outputGrid.Imd.Clear()

	outputGrid.Grid = make([][]uint32, xSteps)
	for i := range outputGrid.Grid {
		outputGrid.Grid[i] = make([]uint32, 12)
	}

	blockWidth := outputGrid.Rect.W() / float64(xSteps)
	blockHeight := outputGrid.Rect.H() / float64(ySteps)

	// Vertical Lines
	for x := 0; x <= xSteps; x++ {
		outputGrid.Imd.Color = ColorGridLines
		outputGrid.Imd.Push(
			pixel.V(
				outputGrid.Rect.Min.X+float64(x)*blockWidth,
				outputGrid.Rect.Min.Y,
			),
			pixel.V(
				outputGrid.Rect.Min.X+float64(x)*blockWidth,
				outputGrid.Rect.Max.Y,
			),
		)
		outputGrid.Imd.Line(1)
	}

	// Horizontal Lines
	for y := 0; y <= ySteps; y++ {
		outputGrid.Imd.Color = ColorGridLines
		outputGrid.Imd.Push(
			pixel.V(
				outputGrid.Rect.Min.X,
				outputGrid.Rect.Min.Y+(float64(y)*blockHeight),
			),
			pixel.V(
				outputGrid.Rect.Max.X,
				outputGrid.Rect.Min.Y+(float64(y)*blockHeight),
			),
		)
		outputGrid.Imd.Line(1)
	}

	outputGrid.isComposing = false
}

func (outputGrid *OutputGrid) Update(win *pixelgl.Window) {

}

func (outputGrid *OutputGrid) DrawTo(imd *imdraw.IMDraw) {
	if !outputGrid.isComposing {
		outputGrid.Imd.Draw(imd)
	}
}
