package estado

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer() *container.Split {

	container1 := container.NewVBox(
		widget.NewLabel("parte 1"),
	)
	container2 := container.NewVBox(
		widget.NewLabel("parte 2"),
	)

	finalContain := container.NewHSplit(
		container1, container2,
	)
	finalContain.SetOffset(0.5)

	return finalContain
}
