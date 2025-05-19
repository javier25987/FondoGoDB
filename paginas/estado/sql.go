package estado

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func getGenData() [8]int {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var datos [8]int
	/*
		0. Total capital
		1. Total aporte a multas
		2. Total multas extra
		3. Total prestamos hechos
		4. Total dinero en prestamos
		5. Total dinero por intereses
		6. Total intereses vencidos (Tabla prestamos hechos)
		7. Total deuda (Tabla prestamos hechos)
	*/

	consulta := "SELECT SUM(ig.capital ), SUM(ig.aporte_a_multas ), SUM(ig.multas_extra ), SUM(p.prestamos_hechos ), SUM(p.dinero_en_prestamos ), SUM(p.dinero_por_intereses ), (SELECT SUM(ph.intereses_vencidos ) FROM prestamos_hechos ph ), (SELECT SUM(ph.deuda ) FROM prestamos_hechos ph ) FROM informacion_general ig JOIN prestamos p ON p.id = ig.id"

	err = db.QueryRow(consulta).Scan(
		&datos[0], &datos[1], &datos[2], &datos[3],
		&datos[4], &datos[5], &datos[6], &datos[7],
	)
	if err != nil {
		log.Fatal(err)
	}

	return datos
}
