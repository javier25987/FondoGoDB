package anotaciones

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("anotaciones"),
	)
}
