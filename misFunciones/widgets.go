package misfunciones

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func Divider() *canvas.Line {
	line := canvas.NewLine(color.Gray{Y: 128})
	line.StrokeWidth = 2

	return line
}
