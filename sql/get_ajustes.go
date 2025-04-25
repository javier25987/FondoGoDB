package sql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetAjusteInt(name string) int {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var value int

	consulta := "SELECT valor_n FROM ajustes WHERE ajuste = '" + name + "'"
	err = db.QueryRow(consulta, name).Scan(&value)

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func GetAjusteStr(name string) string {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var value string

	consulta := "SELECT valor_a FROM ajustes WHERE ajuste = '" + name + "'"
	err = db.QueryRow(consulta, name).Scan(&value)

	if err != nil {
		log.Fatal(err)
	}

	return value
}
