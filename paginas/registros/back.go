package registros

import (
	"database/sql"
	"log"

	myfn "fondo/misFunciones"

	_ "github.com/mattn/go-sqlite3"
)

func getRegistros(mes string) [][4]string {
	allUsers := [][4]string{}
	tArray := [4]string{"Fecha", "Ingresos", "Egresos", "Diferencia"}

	allUsers = append(allUsers, tArray)

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := "SELECT fecha, ingreso, egreso, (ingreso - egreso) FROM registros WHERE STRFTIME('%m', fecha) = '" + mes + "'"

	rows, err := db.Query(consulta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	inicioS := "✅ "
	valorDif := ""

	for rows.Next() {
		var (
			fecha      string
			ingreso    int
			egreso     int
			diferencia int
		)

		err := rows.Scan(&fecha, &ingreso, &egreso, &diferencia)
		if err != nil {
			log.Fatal(err)
		}

		if diferencia < 0 {
			inicioS = "⛔️ -"
			diferencia *= -1
		} else {
			inicioS = "✅ "
		}

		valorDif = inicioS + myfn.FormatComas(diferencia)

		tArray[0] = fecha
		tArray[1] = myfn.FormatComas(ingreso)
		tArray[2] = myfn.FormatComas(egreso)
		tArray[3] = valorDif

		allUsers = append(allUsers, tArray)
	}
	return allUsers
}
