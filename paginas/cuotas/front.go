package cuotas

import (
	"fondo/buscar"
	"fondo/globals"
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
)

func MainContainer() *container.Split {
	index := globals.Index
	win := globals.MyWindow

	data := getUserTable(index)
	table := myfn.MakeTableCuotas(data)

	allContainer := container.NewHSplit(
		container.NewScroll(table),
		container.NewScroll(
			container.NewPadded(
				container.NewVBox(
					makeName(index),         // nombre
					makeFormPay(index, win), // formulario de pago
					buscar.SearchContain(),
					makeOpenFile(),
					makeFormBlock(index, win), // bloqueo
				),
			),
		),
	)
	allContainer.SetOffset(0.4)

	return allContainer
}
