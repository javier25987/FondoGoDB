package buscar

import (
	"fyne.io/fyne/v2/widget"

	"fondo/funcs"
	"fondo/global"
)

func SearchContain() *widget.Card {

	numero := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Numero:", Widget: numero, HintText: "Numero del usuario"},
		},
		SubmitText: "Buscar",
		OnSubmit: func() {
			err, numeroUser := funcs.RectNumber(numero.Text)

			if err {
				global.Index = numeroUser
				global.Refresh()
			}
		},
	}

	return widget.NewCard("Buscar Usuario:", "", form)

}
