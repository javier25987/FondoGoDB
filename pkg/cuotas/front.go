package cuotas

import (
	"fondo/buscar"
	"fondo/funcs"
	"fondo/global"

	"fyne.io/fyne/v2/container"
)

func MainContainer() *container.Split {
	index := global.Index
	win := global.MyWindow

	data := getUserTable(index)
	table := funcs.MakeTableCuotas(data)

	allContainer := container.NewHSplit(
		container.NewScroll(table),
		container.NewScroll(
			container.NewPadded(
				container.NewVBox(
					buscar.SearchContain(),
					makeName(index),         // nombre
					makeFormPay(index, win), // formulario de pago
					makeOpenFile(),
					makeFormBlock(index, win), // bloqueo
				),
			),
		),
	)
	allContainer.SetOffset(0.4)

	return allContainer
}
