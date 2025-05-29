package estado

import (
	"fondo/buscar"

	"fyne.io/fyne/v2/container"
)

func MainContainer() *container.Split {

	container1 := container.NewVScroll(
		makeGenEstado(),
	)
	container2 := container.NewVScroll(
		buscar.SearchContain(),
	)

	finalContain := container.NewHSplit(
		container1, container2,
	)
	finalContain.SetOffset(0.5)

	return finalContain
}
