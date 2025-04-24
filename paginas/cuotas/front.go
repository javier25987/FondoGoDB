package cuotas

import (
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
)

func MainContainer(index int) *container.Split {
	data := getUserTable(index)

	table := myfn.MakeTableCuotas(data)

	nombre := createName(index)

	allContainer := container.NewHSplit(
		container.NewHScroll(table),
		container.NewVBox(
			nombre,
		),
	)
	allContainer.SetOffset(0.35)

	return allContainer
}
