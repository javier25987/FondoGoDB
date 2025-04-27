package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetValueStr(tabla string, columna string, id int, value string) {

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := fmt.Sprintf(
		"UPDATE %s SET %s = '%s' WHERE id = %d",
		tabla, columna, value, id,
	)

	_, err = db.Exec(consulta)

	if err != nil {
		log.Fatal(err)
	}
}

func SetValueInt(tabla string, columna string, id int, value int) {

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := fmt.Sprintf(
		"UPDATE %s SET %s = %d WHERE id = %d",
		tabla, columna, value, id,
	)

	_, err = db.Exec(consulta)

	if err != nil {
		log.Fatal(err)
	}
}
