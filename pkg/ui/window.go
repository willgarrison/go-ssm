package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewWindow(rect pixel.Rect) *pixelgl.Window {

	config := pixelgl.WindowConfig{
		Title:     "Pixel",
		Bounds:    rect,
		Resizable: false,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	return win
}
