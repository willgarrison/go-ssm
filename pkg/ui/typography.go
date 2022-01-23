package ui

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gomono"
)

type Typography struct {
	TxtBatch *pixel.Batch
	Txt      *text.Text
}

func NewTypography(fontSize float64) *Typography {

	typ := new(Typography)

	// Go Font
	ttf, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(ttf, &truetype.Options{Size: fontSize})

	// Custom Font
	// face, err := loadTTF("./fonts/RobotoMono-Bold.ttf", fontSize)
	// if err != nil {
	// 	panic(err)
	// }

	txtAtlas := text.NewAtlas(face, text.ASCII)

	typ.TxtBatch = pixel.NewBatch(&pixel.TrianglesData{}, txtAtlas.Picture())
	typ.Txt = text.New(pixel.ZV, txtAtlas)

	return typ
}

func (typ *Typography) DrawRuneToBatch(r rune, vec pixel.Vec, clr color.Color, txtBatch *pixel.Batch, txt *text.Text) {
	txt.Clear()
	txt.Color = clr
	txt.Dot = vec
	txt.WriteRune(r)
	txt.Draw(txtBatch, pixel.IM)
}

func (typ *Typography) DrawStringToBatch(s string, vec pixel.Vec, clr color.Color, txtBatch *pixel.Batch, txt *text.Text) {
	txt.Clear()
	txt.Color = clr
	txt.Dot = vec
	txt.WriteString(s)
	txt.Draw(txtBatch, pixel.IM)
}

// func loadTTF(path string, size float64) (font.Face, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	bytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	font, err := truetype.Parse(bytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return truetype.NewFace(font, &truetype.Options{
// 		Size:              size,
// 		GlyphCacheEntries: 1,
// 	}), nil
// }
