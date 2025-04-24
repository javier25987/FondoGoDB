package sql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetValueInt(table string, colum string, id int) int {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var value int

	consulta := "SELECT " + colum + " FROM " + table + " WHERE id = ?"
	err = db.QueryRow(consulta, id).Scan(&value)

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func GetValueStr(table string, colum string, id int) string {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var value string

	consulta := "SELECT " + colum + " FROM " + table + " WHERE id = ?"
	err = db.QueryRow(consulta, id).Scan(&value)

	if err != nil {
		log.Fatal(err)
	}

	return value
}
