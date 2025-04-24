package cuotas

import (
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer(index int) *container.Split {
	data := getUserTable(index)

	table := myfn.MakeTableCuotas(data)

	allContainer := container.NewHSplit(
		container.NewHScroll(table),
		container.NewVBox(
			widget.NewLabel("Contenido adicional"),
			widget.NewLabel("Más información"),
			widget.NewLabel("Detalles"),
		),
	)
	allContainer.SetOffset(0.35)

	return allContainer
}
