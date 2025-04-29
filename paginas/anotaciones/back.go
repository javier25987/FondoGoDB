package anotaciones

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
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

	for rows.Next() {
		var (
			general   string
			monetaria string
			multa     string
			acuerdo   string
		)

		err := rows.Scan(&general, &monetaria, &multa, &acuerdo)
		if err != nil {
			log.Fatal(err)
		}

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
	}

	contain.Refresh()
}
