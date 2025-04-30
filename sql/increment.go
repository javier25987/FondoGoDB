package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func IncrementInt(tabla string, columna string, id int, value int) {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := fmt.Sprintf(
		"UPDATE %s SET %s = %s + %d WHERE id = %d",
		tabla, columna, columna, value, id,
	)

	_, err = db.Exec(consulta)

	if err != nil {
		log.Fatal(err)
	}
}

func IncremenStr(tabla string, columna string, id int, value string) {
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var valueDb string

	consulta := "SELECT " + columna + " FROM " + tabla + " WHERE id = ?"
	err = db.QueryRow(consulta, id).Scan(&valueDb)

	if err != nil {
		log.Fatal(err)
	}

	/*
		Aca ya hemos cargado en valor extraido de la base y lo vamos a cargar
		el valor en newValue
	*/

	if valueDb == "n" {
		valueDb = value
	} else {
		valueDb += "_" + value
	}

	consulta = fmt.Sprintf(
		"UPDATE %s SET %s = '%s' WHERE id = %d",
		tabla, columna, valueDb, id,
	)

	_, err = db.Exec(consulta)

	if err != nil {
		log.Fatal(err)
	}
}
