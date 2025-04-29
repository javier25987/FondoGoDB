package anotaciones

import (
	"fondo/globals"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer() *container.Split {

	Index := globals.Index

	contNotas := container.NewVBox()
	cargarAnotaciones(Index, contNotas)

	container1 := container.NewScroll(contNotas)

	container2 := container.NewVBox(
		widget.NewLabel("parte 2"),
	)

	finalContain := container.NewHSplit(
		container1, container2,
	)
	finalContain.SetOffset(0.5)

	return finalContain
}
