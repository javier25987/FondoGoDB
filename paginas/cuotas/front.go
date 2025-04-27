package cuotas

import (
	"fondo/globals"
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
)

func MainContainer() *container.Split {
	index := globals.Index
	win := globals.MyWindow

	data := getUserTable(index)
	table := myfn.MakeTableCuotas(data)

	// todos estos son los elementos habituales pero los planeo manejar por tarjetas
	nombre := makeName(index)
	formulario := makeFormPay(index, win)
	desBloqueo := makeFormBlock(index, win)

	allContainer := container.NewHSplit(
		container.NewHScroll(table),
		container.NewPadded(
			container.NewVBox(
				nombre,
				formulario,
				desBloqueo,
			),
		),
	)
	allContainer.SetOffset(0.43)

	return allContainer
}
