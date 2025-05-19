package estado

import (
	"fmt"

	"fyne.io/fyne/v2/widget"

	myfn "fondo/misFunciones"
)

func makeGenEstado() *widget.Card {
	datos := getGenData()

	nombres := [9]string{
		"Total capital:",
		"Total multas pagas:",
		"Total multas temporales:",
		"Total de prestamos hechos:",
		"Total dinero retirado en prestamos:",
		"Total intereses pagados:",
		"Total Intereses en deuda:",
		"Total deuda por prestamos:",
		"Total pagado en transferencias:",
	}

	var texto string
	var info string

	for i, dato := range datos {
		info = fmt.Sprintf("* **%s** _%s_\n", nombres[i], myfn.FormatComas(dato))
		texto += info
	}

	return widget.NewCard(
		"Estado general:", "",
		widget.NewRichTextFromMarkdown(texto),
	)
}
