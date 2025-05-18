package buscar

import (
	"fyne.io/fyne/v2/widget"

	"fondo/globals"
	myfn "fondo/misFunciones"
)

func SearchContain() *widget.Card {

	numero := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Numero:", Widget: numero, HintText: "Numero del usuario"},
		},
		SubmitText: "Buscar",
		OnSubmit: func() {
			err, numeroUser := myfn.RectNumber(numero.Text)

			if err {
				globals.Index = numeroUser
				globals.Refresh()
			}
		},
	}

	return widget.NewCard("Buscar Usuario:", "", form)

}
