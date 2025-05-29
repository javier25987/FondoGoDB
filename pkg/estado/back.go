package estado

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"fondo/funcs"
)

func makeGenEstado() *fyne.Container {
	contain := container.NewVBox()

	datos := getGenData()

	nombres := [9]string{
		"Total capital:",
		"Total multas pagas:",
		"Total anotaciones:",
		"Total de prestamos hechos:",
		"Total dinero retirado en prestamos:",
		"Total intereses pagados:",
		"Total Intereses en deuda:",
		"Total deuda de todos los prestamos:",
		"Total pagado en transferencias:",
	}

	var info string

	contain.Add(widget.NewCard("Estado general:", "", nil))

	for i, dato := range datos {
		info = fmt.Sprintf("**%s** \n\n _%s_", nombres[i], funcs.FormatComas(dato))
		contain.Add(
			widget.NewCard(
				"", "", widget.NewRichTextFromMarkdown(info),
			),
		)
	}

	return contain
}
