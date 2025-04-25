package cuotas

import (
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
)

func MainContainer(index int) *container.Split {
	data := getUserTable(index)

	table := myfn.MakeTableCuotas(data)

	// todos estos son los elementos habituales pero los planeo manejar por tarjetas
	nombre := makeName(index)
	formulario := makeFormPay(index)
	desBloqueo := makeFormBlock(index)

	allContainer := container.NewHSplit(
		container.NewHScroll(table),
		container.NewVBox(
			nombre,
			formulario,
			desBloqueo,
		),
	)
	allContainer.SetOffset(0.35)

	return allContainer
}
