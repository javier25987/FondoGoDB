package transferencias

import (
	"fondo/funcs"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var User int = -1

func MainContainer() *container.Split {
	ProporcionTabla := [4]float32{80.0, 200.0, 200.0, 200.0}

	dataContainer := container.NewScroll(
		funcs.MakeTable4(
			getUsersTable(), ProporcionTabla,
		),
	)

	selectUser := widget.NewSelect(
		getAllUsers(), func(s string) {
			if s == "TODOS" {
				dataContainer.Content = funcs.MakeTable4(
					getUsersTable(), ProporcionTabla,
				)
			} else {
				intValue, _ := strconv.Atoi(s)
				User = intValue

				dataContainer.Content = funcs.MakeTable4(
					getUserTable(User), ProporcionTabla,
				)
			}
			dataContainer.Refresh()
		},
	)
	selectUser.PlaceHolder = "TODOS"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:     "Buscar usuario:",
				Widget:   selectUser,
				HintText: "Seleccione el numero del usuario",
			},
		},
	}

	allContainer := container.NewVSplit(
		widget.NewCard("Buscar usuario:", "", form),
		dataContainer,
	)
	allContainer.SetOffset(0.01)

	return allContainer
}
