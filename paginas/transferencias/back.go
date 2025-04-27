package transferencias

import (
	"database/sql"
	"log"
	"strconv"

	myfn "fondo/misFunciones"

	_ "github.com/mattn/go-sqlite3"
)

func getAllUsers() []string {
	allUsers := []string{}
	allUsers = append(allUsers, "TODOS")

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT id FROM transferencias ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		allUsers = append(allUsers, strconv.Itoa(id))
	}
	return allUsers
}

func getUsersTable() [][4]string {
	allUsers := [][4]string{}

	allUsers = append(
		allUsers,
		[4]string{"Numero", "Nombre", "Fecha", "Monto"},
	)

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT t.id, ig.nombre, t.fecha, t.monto FROM transferencias t JOIN informacion_general ig ON t.id = ig.id",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id       int
		nombre   string
		fecha    string
		monto    int
		fmtMonto string
	)

	for rows.Next() {

		err := rows.Scan(&id, &nombre, &fecha, &monto)
		if err != nil {
			log.Fatal(err)
		}

		fmtMonto = myfn.FormatComas(monto)

		allUsers = append(
			allUsers,
			[4]string{strconv.Itoa(id), nombre, fecha, fmtMonto},
		)
	}
	return allUsers
}

func getUserTable(index int) [][4]string {
	allUsers := [][4]string{}

	allUsers = append(
		allUsers,
		[4]string{"Numero", "Nombre", "Fecha", "Monto"},
	)

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT t.id, ig.nombre, t.fecha, t.monto FROM transferencias t JOIN informacion_general ig ON t.id = ig.id WHERE t.id = ?",
		index,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id       int
		nombre   string
		fecha    string
		monto    int
		fmtMonto string
	)

	for rows.Next() {

		err := rows.Scan(&id, &nombre, &fecha, &monto)
		if err != nil {
			log.Fatal(err)
		}

		fmtMonto = myfn.FormatComas(monto)

		allUsers = append(
			allUsers,
			[4]string{strconv.Itoa(id), nombre, fecha, fmtMonto},
		)
	}

	return allUsers
}
