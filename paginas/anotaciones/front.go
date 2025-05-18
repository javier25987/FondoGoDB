package anotaciones

import (
	"fondo/buscar"
	"fondo/globals"

	"fyne.io/fyne/v2/container"
)

func MainContainer() *container.Split {

	Index := globals.Index

	contNotas := container.NewVBox()
	cargarAnotaciones(Index, contNotas)

	container1 := container.NewScroll(contNotas)

	container2 := container.NewVBox(
		makeName(Index),        // nombre
		makeHacerMulta(Index),  //formato para hacer una anotacion
		buscar.SearchContain(), // buscar usuario
		makePago(Index),        // hacer un pago
	)

	finalContain := container.NewHSplit(
		container1, container2,
	)
	finalContain.SetOffset(0.5)

	return finalContain
}
