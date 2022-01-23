package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type SSM struct {
	Imd         *imdraw.IMDraw
	Text        *Typography
	Rect        pixel.Rect
	isComposing bool
	Pattern     []rune
	Grid        [][]uint32
}

func NewSSM(rect pixel.Rect, pattern []rune) *SSM {

	ssm := &SSM{
		Imd:     imdraw.New(nil),
		Text:    NewTypography(11),
		Rect:    rect,
		Pattern: pattern,
	}

	ssm.Clear()
	ssm.Compose()

	return ssm
}

func (ssm *SSM) Clear() {
	ssm.Imd.Clear()
	ssm.Text.TxtBatch.Clear()
}

func (ssm *SSM) Compose() {

	ssm.isComposing = true

	ssm.Grid = make([][]uint32, len(ssm.Pattern))
	for i := range ssm.Grid {
		ssm.Grid[i] = make([]uint32, len(ssm.Pattern))
	}

	blockW := ssm.Rect.W() / float64(len(ssm.Pattern))
	blockH := ssm.Rect.H() / float64(len(ssm.Pattern))

	xPattern := ssm.Pattern
	yPattern := ssm.Pattern

	// Draw active blocks
	for x := range ssm.Pattern {
		for y := range ssm.Pattern {

			ssm.Imd.Color = ColorBlockSystemOn

			// the diagonal
			if x == y {
				ssm.Imd.Color = ColorBlockSystemOnDiagonal
			}

			// spaces
			if xPattern[x] == ' ' {
				ssm.Imd.Color = ColorBlockSystemOnSpace
			}

			// if unique
			if unique(xPattern[x], xPattern) {
				ssm.Imd.Color = ColorBlockSystemOnUnique
			}

			if xPattern[x] == yPattern[y] {
				ssm.Imd.Push(
					pixel.V(
						ssm.Rect.Min.X+(float64(x)*blockW),
						ssm.Rect.Max.Y-(float64(y)*blockH),
					),
					pixel.V(
						ssm.Rect.Min.X+(float64(x)*blockW)+blockW,
						ssm.Rect.Max.Y-(float64(y)*blockH)-blockH,
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
				ssm.Rect.Min.X+float64(x)*blockW,
				ssm.Rect.Min.Y,
			),
			pixel.V(
				ssm.Rect.Min.X+float64(x)*blockW,
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
				ssm.Rect.Min.Y+(float64(y)*blockH),
			),
			pixel.V(
				ssm.Rect.Max.X,
				ssm.Rect.Min.Y+(float64(y)*blockH),
			),
		)
		ssm.Imd.Line(1)
	}

	// Draw Text
	for x := range ssm.Grid {
		str := ssm.Pattern[x]
		strX := ssm.Rect.Min.X + float64(x)*blockW + blockW/2
		ssm.Text.DrawRuneToBatch(str, pixel.V(strX, 30), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
		ssm.Text.DrawRuneToBatch(str, pixel.V(strX, ssm.Rect.Max.Y+20), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
	}
	for y := range ssm.Grid {
		str := ssm.Pattern[y]
		strY := ssm.Rect.Max.Y - (float64(y) * blockH) - blockH/2
		ssm.Text.DrawRuneToBatch(str, pixel.V(30, strY), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
		ssm.Text.DrawRuneToBatch(str, pixel.V(ssm.Rect.Max.X+20, strY), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
	}

	if len(ssm.Pattern) == 0 {
		ssm.Text.DrawStringToBatch("Instructions:", pixel.V(60, ssm.Rect.Max.Y), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
		ssm.Text.DrawStringToBatch("Enter letters, numbers and spaces", pixel.V(60, ssm.Rect.Max.Y-20), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
		ssm.Text.DrawStringToBatch("Escape clears all input", pixel.V(60, ssm.Rect.Max.Y-40), ColorText, ssm.Text.TxtBatch, ssm.Text.Txt)
	}

	ssm.isComposing = false
}

func (ssm *SSM) Update(newPattern []rune) {
	ssm.Pattern = newPattern
	ssm.Clear()
	ssm.Compose()
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

func unique(needle rune, haystack []rune) bool {
	count := int32(0)
	for i := range haystack {
		if haystack[i] == needle {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}
