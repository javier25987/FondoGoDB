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

	anotacion := widget.NewEntry()

	monto := widget.NewEntry()

	motivoAnota := widget.NewRadioGroup(
		[]string{"General", "Monetaria", "Multa", "Acuerdo"},
		func(s string) {

		},
	)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Anotacion:", Widget: anotacion, HintText: "No incluya los simbolos { _ [ ] }"},
			{Text: "Monto:", Widget: monto, HintText: "Se pueden incluir numeros negativos"},
			{Text: "Motivo:", Widget: motivoAnota},
		},
		SubmitText: "Realizar anotacion",
		OnSubmit: func() {

		},
	}

	nombre := makeName(Index)

	container2 := container.NewVBox(
		nombre,
		widget.NewCard(
			"Hacer una anotacion:", "", form,
		),
	)

	finalContain := container.NewHSplit(
		container1, container2,
	)
	finalContain.SetOffset(0.6)

	return finalContain
}
