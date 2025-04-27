package cuotas

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"fondo/globals"
	myfn "fondo/misFunciones"
	mySQL "fondo/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	pagas := mySQL.GetValueInt("cuotas", "pagas", index)
	adeudas := mySQL.GetValueInt("cuotas", "adeudas", index)
	multas := mySQL.GetValueStr("cuotas", "multas", index)
	bloqueos := mySQL.GetValueStr("cuotas", "bloqueos", index)

	// creamos los arrays para codificar los datos
	semanas_array := [50]string{}
	pagas_array := [50]string{}

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
	nombre = strings.Title(nombre)
	puestos := mySQL.GetValueInt("informacion_general", "puestos", index)
	mensaje := fmt.Sprintf("№ %d - %s : %d puesto(s)", index, nombre, puestos)

	telefono := mySQL.GetValueStr("informacion_general", "telefono", index)
	numeroT := fmt.Sprintf("Numero de telefono: %s", telefono)

	// widget.NewRichTextFromMarkdown(mensaje),

	return container.NewVBox(
		widget.NewCard(mensaje, numeroT, nil),
	)
}

func makeFormPay(index int, win *fyne.Window) fyne.CanvasObject {

	var (
		cuotasAPagar int    = 0
		multasAPagar int    = 0
		metodoDePago string = ""
	)

	cuotasPagar := widget.NewSelect(
		myfn.MakeRange(0, 10),
		func(s string) {
			cuotasAPagar, _ = strconv.Atoi(s)
		},
	)
	cuotasPagar.PlaceHolder = "0"

	multasPagar := widget.NewSelect(
		myfn.MakeRange(0, 5),
		func(s string) {
			multasAPagar, _ = strconv.Atoi(s)
		},
	)
	multasPagar.PlaceHolder = "0"

	metodoPago := widget.NewRadioGroup(
		[]string{"Efectivo", "Transferencia"},
		func(s string) {
			metodoDePago = s
		},
	)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Cuotas a pagar:", Widget: cuotasPagar}, // HintText: "Your full name"},
			{Text: "Multas a pagar:", Widget: multasPagar},
			{Text: "Metodo de pago:", Widget: metodoPago},
		},
		SubmitText: "Pagar",
		OnSubmit: func() {
			result, err := rectificarPago(cuotasAPagar, multasAPagar, metodoDePago)

			if result {
				pagarCuotas(index, cuotasAPagar, multasAPagar, metodoDePago)
				globals.Refresh()
			} else {
				dialog.ShowError(errors.New(err), *win)
			}
		},
	}

	return widget.NewCard("Formulario de pago:", "", form)
}

func makeFormBlock(index int, win *fyne.Window) fyne.CanvasObject {

	semanasBlock := widget.NewEntry()
	semanasBlock.PlaceHolder = "1 - 50"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:     "Semana a (des)bloquear:",
				Widget:   semanasBlock,
				HintText: "Escriba un numero entre 1 y 50",
			},
		},
		SubmitText: "(Des)Bloquear",
		OnSubmit: func() {
			result, err := rectificarBloqueo(semanasBlock.Text)

			if result {
				makeBlock(index, semanasBlock.Text)
				globals.Refresh()
			} else {
				dialog.ShowError(errors.New(err), *win)
			}
		},
	}
	return widget.NewCard("(Des)Bloquear semanas:", "", form)
}

func makeOpenFile() fyne.CanvasObject {

	contButtonOpen := container.NewCenter(
		widget.NewButton(
			"Abrir cheque", func() {

			},
		),
	)

	return widget.NewCard("Abrir ultimo cheque:", "", contButtonOpen)
}
