package registros

import (
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer() *container.Split {

	proporcionTable := [4]float32{100, 150, 150, 150}

	container2 := container.NewScroll(
		myfn.MakeTable4(getRegistros("01"), proporcionTable),
	)

	mesesSelect := widget.NewSelect(
		[]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"},
		func(s string) {
			container2.Content = myfn.MakeTable4(
				getRegistros(s), proporcionTable,
			)
			container2.Refresh()
		},
	)
	mesesSelect.PlaceHolder = "01"

	formRegistro := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:     "Mes a buscar:",
				Widget:   mesesSelect,
				HintText: "Introduzca el mes a buscar",
			},
		},
	}

	container1 := container.NewVBox(
		widget.NewCard("Buscar por mes:", "", formRegistro),
	)

	finalContain := container.NewVSplit(
		container1, container2,
	)
	finalContain.SetOffset(0.01)

	return finalContain
}
