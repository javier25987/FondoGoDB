package cuotas

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	myfn "fondo/misFunciones"
	mySQL "fondo/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

	bloqueos := "9_7_2_19"

	if bloqueos != "n" {
		bloqueosArray := strings.Split(bloqueos, "_")

		var b_ind int

		for _, b := range bloqueosArray {
			b_ind, _ = strconv.Atoi(b) // descartar el error es una mala practica pero no me importa
			pagas_array[b_ind-1] = "b"
		}
	}

	for i := 0; i < 50; i++ {
		semanas_array[i] = strconv.Itoa(i + 1)

		if pagas_array[i] == "b" {
			pagas_array[i] = "🔒"
			continue
		}

		if pagas > 0 {
			pagas_array[i] = "✅"
			pagas--
		} else if adeudas > 0 {
			pagas_array[i] = "❌"
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

	// creamos el calendario de pagos

	calendario := mySQL.GetAjusteStr("calendario")
	calendarioArray := strings.Split(calendario, "_")

	for i, f := range calendarioArray {
		calendarioArray[i] = f[:len(f)-3]
	}

	// llenamos el array con los datos
	for i := 0; i < 50; i++ {
		myArray[i+1][0] = semanas_array[i]
		myArray[i+1][1] = calendarioArray[i]
		myArray[i+1][2] = pagas_array[i]
		myArray[i+1][3] = multas_array[i]
	}

	return myArray
}

func makeName(index int) *fyne.Container {

	nombre := mySQL.GetValueStr("informacion_general", "nombre", index)
	nombre = myfn.Title(nombre)

	puestos := mySQL.GetValueInt("informacion_general", "puestos", index)

	mensaje := fmt.Sprintf("# № %d - %s : %d puesto(s)", index, nombre, puestos)

	// return container.NewVBox(
	//     widget.NewRichTextFromMarkdown(mensaje),
	// )

	return container.NewVBox(
		widget.NewCard(mensaje, "", nil),
		// widget.NewRadioGroup([]string{"6:30pm", "7:00pm", "7:45pm"}, func(string) {}),
	)
}

func makeFormPay(index int) fyne.CanvasObject {

	cuotasPagar := widget.NewSelect(
		[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		func(s string) {},
	)
	cuotasPagar.PlaceHolder = "0"

	multasPagar := widget.NewSelect(
		[]string{"0", "1", "2"},
		func(s string) {},
	)
	multasPagar.PlaceHolder = "0"

	metodoPago := widget.NewRadioGroup(
		[]string{"Efectivo", "Transferencia"},
		func(s string) {},
	)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Cuotas a pagar:", Widget: cuotasPagar}, // HintText: "Your full name"},
			{Text: "Multas a pagar:", Widget: multasPagar},
			{Text: "Metodo de pago:", Widget: metodoPago},
		},
		SubmitText: "Pagar",
		OnSubmit:   func() {},
	}

	return widget.NewCard("Formulario de pago:", "", form)
}
