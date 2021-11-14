package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/willgarrison/go-iobox"
)

type SSM struct {
	IOBox       *iobox.IOBox
	Signal      iobox.Signal
	Imd         *imdraw.IMDraw
	Text        *Typography
	Rect        pixel.Rect
	lastMouseX  int32
	lastMouseY  int32
	isComposing bool
	Pattern     []string
	Grid        [][]uint32
}

func NewSSM(rect pixel.Rect, pattern []string) *SSM {

	ssm := &SSM{
		IOBox:      iobox.New(),
		Imd:        imdraw.New(nil),
		Text:       NewTypography(12),
		Rect:       rect,
		lastMouseX: -1,
		lastMouseY: -1,
		Pattern:    pattern,
	}

	ssm.Compose()

	return ssm
}

func (ssm *SSM) Compose() {

	ssm.isComposing = true

	ssm.Imd.Clear()

	ssm.Grid = make([][]uint32, len(ssm.Pattern))
	for i := range ssm.Grid {
		ssm.Grid[i] = make([]uint32, len(ssm.Pattern))
	}

	blockSize := ssm.Rect.W() / float64(len(ssm.Pattern))

	xPattern := ssm.Pattern
	yPattern := ssm.Pattern
	padding := 0.0

	// Draw active blocks
	for x := range ssm.Pattern {
		for y := range ssm.Pattern {

			if xPattern[x] == "-" {
				continue
			}

			ssm.Imd.Color = ColorBlockSystemOn
			padding = 0.0

			// the diagonal
			if x == y {
				ssm.Imd.Color = ColorBlockSystemOnDiagonal
				// padding = 7.0
			}

			// if unique
			if occurrence(xPattern[x], xPattern) == 1 {
				ssm.Imd.Color = ColorBlockSystemOnUnique
				padding = -7.0
			}

			if xPattern[x] == yPattern[y] {
				ssm.Imd.Push(
					pixel.V(
						ssm.Rect.Min.X+(float64(x)*blockSize)+padding,
						ssm.Rect.Max.Y-(float64(y)*blockSize)-padding,
					),
					pixel.V(
						ssm.Rect.Min.X+(float64(x)*blockSize)+blockSize-padding,
						ssm.Rect.Max.Y-(float64(y)*blockSize)-blockSize+padding,
					),
				)

				ssm.Imd.Rectangle(0)
			}
		}
	}

	// Vertical Lines
	for x := 0; x <= len(ssm.Pattern); x++ {
		ssm.Imd.Color = ColorGridLines
		ssm.Imd.Push(
			pixel.V(
				ssm.Rect.Min.X+float64(x)*blockSize,
				ssm.Rect.Min.Y,
			),
			pixel.V(
				ssm.Rect.Min.X+float64(x)*blockSize,
				ssm.Rect.Max.Y,
			),
		)
		ssm.Imd.Line(1)
	}

	// Horizontal Lines
	for y := 0; y <= len(ssm.Pattern); y++ {
		ssm.Imd.Color = ColorGridLines
		ssm.Imd.Push(
			pixel.V(
				ssm.Rect.Min.X,
				ssm.Rect.Min.Y+(float64(y)*blockSize),
			),
			pixel.V(
				ssm.Rect.Max.X,
				ssm.Rect.Min.Y+(float64(y)*blockSize),
			),
		)
		ssm.Imd.Line(1)
	}

	// Draw Text
	ssm.Text.TxtBatch.Clear()
	for x := range ssm.Grid {
		str := ssm.Pattern[x]
		strX := ssm.Rect.Min.X + float64(x)*blockSize + blockSize/2 - (ssm.Text.Txt.BoundsOf(str).W() / 2)
		ssm.Text.DrawTextToBatch(str, pixel.V(strX, 30), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
		ssm.Text.DrawTextToBatch(str, pixel.V(strX, ssm.Rect.Max.Y+20), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
	}
	for y := range ssm.Grid {
		str := ssm.Pattern[y]
		strY := ssm.Rect.Max.Y - (float64(y) * blockSize) - blockSize/2 - (ssm.Text.Txt.BoundsOf(str).H() / 2)
		ssm.Text.DrawTextToBatch(str, pixel.V(30, strY), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
		ssm.Text.DrawTextToBatch(str, pixel.V(ssm.Rect.Max.X+20, strY), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
	}

	ssm.isComposing = false
}

func (ssm *SSM) Update(win *pixelgl.Window) {

}

func (ssm *SSM) DrawTo(imd *imdraw.IMDraw) {
	if !ssm.isComposing {
		ssm.Imd.Draw(imd)
	}
}

func (ssm *SSM) DrawTextTo(win *pixelgl.Window) {
	if !ssm.isComposing {
		ssm.Text.TxtBatch.Draw(win)
	}
}

func occurrence(needle string, haystack []string) int32 {
	count := int32(0)
	for i := range haystack {
		if haystack[i] == needle {
			count++
		}
	}
	return count
}
