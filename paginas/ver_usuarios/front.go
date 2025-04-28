package ver_usuarios

import (
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer() *container.Split {

	porTablaBuscar := [4]float32{70, 200, 100, 100}

	containTablas := container.NewScroll(
		myfn.MakeTable4(
			getAllUsers(), porTablaBuscar,
		),
	)

	nombreABuscar := widget.NewEntry()

	formBuscarNombre := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:     "Nombre a buscar:",
				Widget:   nombreABuscar,
				HintText: "Escriba el nombre o un fracmento de este para buscar",
			},
		},
		SubmitText: "Buscar",
		OnSubmit: func() {
			if nombreABuscar.Text == "" {
				containTablas.Content = myfn.MakeTable4(
					getAllUsers(), porTablaBuscar,
				)
			} else {
				containTablas.Content = myfn.MakeTable4(
					getNameUsers(nombreABuscar.Text), porTablaBuscar,
				)
			}
			containTablas.Refresh()
		},
	}

	containAcuerdos := container.NewCenter(
		widget.NewButton(
			"Consultar", func() {
				containTablas.Content = myfn.MakeTable4(
					getAcuerdos(), porTablaBuscar,
				)
				containTablas.Refresh()
			},
		),
	)

	containTarget := container.NewScroll(
		container.NewVBox(
			widget.NewCard("Buscar por nombre:", "", formBuscarNombre),
			widget.NewCard(
				"Consultar acuerdo:",
				"Aca se muestra a todos los que tengan que firmar acuerdo",
				containAcuerdos,
			),
			widget.NewCard("Buscar boletas:", "Pendiente a actualizacion", nil),
		),
	)

	containS := container.NewHSplit(
		containTablas, containTarget,
	)
	containS.SetOffset(0.5)

	return containS
}
