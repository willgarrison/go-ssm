package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewWindow(title string, rect pixel.Rect) *pixelgl.Window {

	config := pixelgl.WindowConfig{
		Title:     title,
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
