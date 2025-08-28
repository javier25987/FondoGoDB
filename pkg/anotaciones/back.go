package anotaciones

import (
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	mySQl "fondo/sql"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// La funcion cargarcargarAnotaciones() lo unico que hace es tomar un contenedor y cargar en el
// una lista de todas las anotaciones realizadas, esta funcion toma como argumentos el index del
// usuario para el cual queremos leer sus anotaciones y el contenedor al cual lo vamos a cargar
// todo
func cargarAnotaciones(index int, contain *fyne.Container) {

	// limpiar contanido del contenedor
	contain.Objects = []fyne.CanvasObject{}

	// leer todas las cargarAnotaciones
	/*
		Este metodo puede ser un poco menos eficiente ya que hace 4 consultas por separado
		en ves de una sola pero me parece mas practico ya que usa la capa de abstraccion
		creada con anteriooridad
	*/
	general := mySQl.GetValueStr("anotaciones", "general", index)
	monetaria := mySQl.GetValueStr("anotaciones", "monetaria", index)
	multa := mySQl.GetValueStr("anotaciones", "multa", index)
	acuerdo := mySQl.GetValueStr("anotaciones", "acuerdo", index)

	mensage := "# Anotaciones hechas:\n"

	if general != "n" {
		mensage += "## Generales:\n"
		for i := range strings.SplitSeq(general, "_") {
			mensage += "* " + i + "\n"
		}
	}

	if monetaria != "n" {
		mensage += "## Monetarias:\n"
		for i := range strings.SplitSeq(monetaria, "_") {
			mensage += "* " + i + "\n"
		}
	}

	if multa != "n" {
		mensage += "## Multas:\n"
		for i := range strings.SplitSeq(multa, "_") {
			mensage += "* " + i + "\n"
		}
	}

	if acuerdo != "n" {
		mensage += "## Acuerdos:\n"
		for i := range strings.SplitSeq(acuerdo, "_") {
			mensage += "* " + i + "\n"
		}
	}

	contain.Add(widget.NewRichTextFromMarkdown(mensage))
	contain.Refresh()
}

func makeName(index int) *widget.Card {
	nombre := mySQl.GetValueStr("informacion_general", "nombre", index)
	nombre = strings.Title(nombre)

	mensaje := fmt.Sprintf("№ %d : %s", index, nombre) // widget.NewRichTextFromMarkdown(mensaje),

	return widget.NewCard(mensaje, "", nil)
}

func makeHacerMulta(index int) *widget.Card {
	anotacion := widget.NewEntry()

	monto := widget.NewEntry()

	motivoAnota := widget.NewSelect(
		[]string{"General", "Monetaria", "Multa", "Acuerdo"},
		func(s string) {

		},
	)
	motivoAnota.PlaceHolder = "General"

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
	return widget.NewCard("Hacer una anotacion:", "", form)
}

func makePago(index int) *widget.Card {
	return widget.NewCard("Pagar:", "", nil)
}
