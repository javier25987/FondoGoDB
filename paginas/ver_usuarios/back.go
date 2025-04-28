package ver_usuarios

import (
	"database/sql"
	"log"
	"strconv"

	myfn "fondo/misFunciones"

	_ "github.com/mattn/go-sqlite3"
)

func getAllUsers() [][4]string {
	allUsers := [][4]string{}
	tArray := [4]string{"Numero", "Nombre", "Telefono", "Estado"}

	allUsers = append(allUsers, tArray)

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := "SELECT ig.id, ig.nombre, ig.telefono, ig.estado From informacion_general ig"

	rows, err := db.Query(consulta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id       int
			nombre   string
			telefono string
			estado   bool
		)

		err := rows.Scan(&id, &nombre, &telefono, &estado)
		if err != nil {
			log.Fatal(err)
		}

		tArray[0] = strconv.Itoa(id)
		tArray[1] = nombre
		tArray[2] = telefono

		if estado {
			tArray[3] = "✅"
		} else {
			tArray[3] = "❌"
		}

		allUsers = append(allUsers, tArray)
	}
	return allUsers
}

func getNameUsers(nombre string) [][4]string {
	allUsers := [][4]string{}
	tArray := [4]string{"Numero", "Nombre", "Telefono", "Estado"}

	allUsers = append(allUsers, tArray)

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := "SELECT ig.id, ig.nombre, ig.telefono, ig.estado From informacion_general ig WHERE ig.nombre LIKE '%" + nombre + "%'"

	rows, err := db.Query(consulta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id       int
			nombre   string
			telefono string
			estado   bool
		)

		err := rows.Scan(&id, &nombre, &telefono, &estado)
		if err != nil {
			log.Fatal(err)
		}

		tArray[0] = strconv.Itoa(id)
		tArray[1] = nombre
		tArray[2] = telefono

		if estado {
			tArray[3] = "✅"
		} else {
			tArray[3] = "❌"
		}

		allUsers = append(allUsers, tArray)
	}
	return allUsers
}

func getAcuerdos() [][4]string {
	allUsers := [][4]string{}
	tArray := [4]string{"Numero", "Nombre", "Capital", "Dinero retirado"}

	allUsers = append(allUsers, tArray)

	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consulta := "SELECT ig.id, ig.nombre, ig.capital, p.dinero_por_si_mismo From informacion_general ig JOIN prestamos p ON ig.id = p.id WHERE p.dinero_por_si_mismo < ig.capital/2"

	rows, err := db.Query(consulta)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id      int
			nombre  string
			capital int
			dinero  int
		)

		err := rows.Scan(&id, &nombre, &capital, &dinero)
		if err != nil {
			log.Fatal(err)
		}

		tArray[0] = strconv.Itoa(id)
		tArray[1] = nombre
		tArray[2] = myfn.FormatComas(capital)
		tArray[3] = myfn.FormatComas(dinero)

		allUsers = append(allUsers, tArray)
	}
	return allUsers
}
