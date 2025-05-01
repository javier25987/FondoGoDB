package anotaciones

import (
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	mySQl "fondo/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
La funcion cargarcargarAnotaciones() lo unico que hace es tomar un contenedor y cargar en el
una lista de todas las anotaciones realizadas, esta funcion toma como argumentos el index del
usuario para el cual queremos leer sus anotaciones y el contenedor al cual lo vamos a cargar
todo
*/
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

	contain.Add(widget.NewCard("GENERALES:", "", nil))
	if general != "n" {
		for i := range strings.SplitSeq(general, "_") {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Add(widget.NewCard("MONETARIAS:", "", nil))
	if monetaria != "n" {
		for i := range strings.SplitSeq(monetaria, "_") {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Add(widget.NewCard("MULTAS:", "", nil))
	if multa != "n" {
		for i := range strings.SplitSeq(multa, "_") {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Add(widget.NewCard("ACUERDOS:", "", nil))
	if acuerdo != "n" {
		for i := range strings.SplitSeq(acuerdo, "_") {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Refresh()
}

func makeName(index int) *fyne.Container {

	nombre := mySQl.GetValueStr("informacion_general", "nombre", index)
	nombre = strings.Title(nombre)

	mensaje := fmt.Sprintf("№ %d : %s", index, nombre)

	// widget.NewRichTextFromMarkdown(mensaje),

	return container.NewVBox(
		widget.NewCard(mensaje, "", nil),
	)
}
