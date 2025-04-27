package transferencias

import (
	myfn "fondo/misFunciones"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var User int = -1

func MainContainer() *container.Split {

	dataContainer := container.NewScroll(
		myfn.MakeTableTransf(
			getUsersTable(),
		),
	)

	selectUser := widget.NewSelect(
		getAllUsers(), func(s string) {
			if s == "TODOS" {
				User = -1
			} else {
				intValue, _ := strconv.Atoi(s)
				User = intValue
			}
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
		SubmitText: "Buscar",
		OnSubmit: func() {
			if User < 0 {
				dataContainer.Content = myfn.MakeTableTransf(
					getUsersTable(),
				)
			} else {
				dataContainer.Content = myfn.MakeTableTransf(
					getUserTable(User),
				)
			}
			dataContainer.Refresh()
		},
	}

	allContainer := container.NewVSplit(
		widget.NewCard("Buscar usuario:", "", form),
		dataContainer,
	)
	allContainer.SetOffset(0.01)

	return allContainer
}
