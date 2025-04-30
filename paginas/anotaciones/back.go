package anotaciones

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	mySQl "fondo/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func cargarAnotaciones(index int, contain *fyne.Container) {

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := fmt.Sprintf("SELECT general, monetaria, multa, acuerdo From anotaciones WHERE id = %d", index)

	rows, err := db.Query(consulta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	contain.Objects = []fyne.CanvasObject{}

	general := mySQl.GetValueStr("anotaciones", "general", index)
	monetaria := mySQl.GetValueStr("anotaciones", "monetaria", index)
	multa := mySQl.GetValueStr("anotaciones", "multa", index)
	acuerdo := mySQl.GetValueStr("anotaciones", "acuerdo", index)

	contain.Add(widget.NewCard("GENERALES:", "", nil))
	if general != "n" {
		generalArray := strings.Split(general, "_")

		for _, i := range generalArray {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Add(widget.NewCard("MONETARIAS:", "", nil))
	if monetaria != "n" {
		monetariaArray := strings.Split(monetaria, "_")

		for _, i := range monetariaArray {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Add(widget.NewCard("MULTAS:", "", nil))
	if multa != "n" {
		multaArray := strings.Split(multa, "_")

		for _, i := range multaArray {
			contain.Add(
				widget.NewCard(
					"", i, nil,
				),
			)
		}
	}

	contain.Add(widget.NewCard("ACUERDOS:", "", nil))
	if acuerdo != "n" {
		acuerdoArray := strings.Split(acuerdo, "_")

		for _, i := range acuerdoArray {
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
