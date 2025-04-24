package cuotas

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

/*
Esta funcion devueleve un array que sera usado para llenar la tabla de cuotas
El array tiene 51 filas y 4 columnas, la primera fila es el encabezado
y las siguientes 50 filas son los datos de las cuotas
*/
func getUserTable(index int) [51][4]string {

	// abrimos la base de datos
	db, err := sql.Open("sqlite3", "./Fondo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// creamos la tabla para los usuarios
	myArray := [51][4]string{}
	myArray[0] = [4]string{"Semana", "Fecha", "Estado", "Multas"}

	// obtenemos los valores del usuario
	var (
		pagas   int
		adeudas int
		multas  string
	)
	err = db.QueryRow(
		"SELECT pagas, adeudas, multas FROM cuotas WHERE id = ?",
		index,
	).Scan(&pagas, &adeudas, &multas)

	if err != nil {
		log.Fatal(err)
	}

	// creamos los arrays para codificar los datos

	semanas_array := [50]string{}
	pagas_array := [50]string{}

	for i := 0; i < 50; i++ {
		semanas_array[i] = strconv.Itoa(i + 1)

		if pagas > 0 {
			pagas_array[i] = "✅"
			pagas--
		} else if adeudas > 0 {
			pagas_array[i] = "🚨"
			adeudas--
		} else {
			pagas_array[i] = ""
		}
	}

	// creamos el array de multas
	multas_array := [50]string{""}

	if multas != "n" {
		mSplit1 := strings.Split(multas, "_")
		var ind int

		for _, m := range mSplit1 {
			mSplit2 := strings.Split(m, ":")

			ind, _ = strconv.Atoi(mSplit2[0]) // descartar el error es una mala practica pero no me importa

			multas_array[ind] = mSplit2[1]
		}
	}

	// llenamos el array con los datos
	for i := 0; i < 50; i++ {
		myArray[i+1][0] = semanas_array[i]
		myArray[i+1][1] = "21/05/2004" // fecha de ejemplo
		myArray[i+1][2] = pagas_array[i]
		myArray[i+1][3] = multas_array[i]
	}

	return myArray
}
