package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int // Type alias

type PxCanvasConfig struct {
	DrawingArea     fyne.Size
	CanvasOffset    fyne.Position
	PxRoows, PxCols int
	PxSize          int
}

type State struct {
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
}


func (state *State) SetFilePath(path string) {
	state.FilePath = path
}