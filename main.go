package main

import (
	// importaciones de la biblioteca standart

	"log"

	// importaciones de fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// importaciones de mis paquetes
	"fondo/globals"
	"fondo/paginas/ajustes"
	"fondo/paginas/anotaciones"
	"fondo/paginas/cuotas"
	"fondo/paginas/estado"
	"fondo/paginas/menu"
	"fondo/paginas/modificar_usuarios"
	"fondo/paginas/prestamos"
	"fondo/paginas/registros"
	"fondo/paginas/rifas"
	"fondo/paginas/transferencias"
	"fondo/paginas/ver_usuarios"

	myfn "fondo/misFunciones"
)

func main() {

	log.Println("Iniciando el programa...")

	myApp := app.New()
	myWindow := myApp.NewWindow("Fondo San Javier")
	globals.MyWindow = &myWindow
	globals.WinDialog = myWindow

	Inyect()

	// Panel principal que irá cambiando con la navegación
	myContainer := container.NewPadded()
	globals.Container2 = myContainer

	// Sidebar
	sidebar := container.NewVScroll(
		make_slide_bar(myWindow),
	)

	entradaUser := widget.NewEntry()

	botonBuscar := widget.NewButton(
		"🔎 Buscar", func() {
			err, numeroUser := myfn.RectNumber(entradaUser.Text)

			if err {
				globals.Index = numeroUser
				globals.Refresh()
			}
		},
	)

	setUser := container.NewVBox(
		widget.NewLabelWithStyle(
			"Buscar Usuarios",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		entradaUser,
		botonBuscar,
	)

	barContain := container.NewVSplit(
		sidebar, setUser,
	)
	barContain.SetOffset(0.9)

	// Contenedor general: Sidebar + Contenido principal
	mainContainer := container.NewHSplit(barContain, myContainer)
	mainContainer.SetOffset(0.18) // Tamaño relativo de la barra lateral

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

/*
Inyect es una funcion que carga todas las funciones principales de cada container a una
estructura para ser usadas por la funcion globals.Refresh() esto mediante el paquete main
evita una importacion circular, esta funcion debe ser usada en la funcion main al inicio.
*/
func Inyect() {
	globals.FuncionesInyect.Menu = menu.MainContainer
	globals.FuncionesInyect.Cuotas = cuotas.MainContainer
	globals.FuncionesInyect.Prestamos = prestamos.MainContainer
	globals.FuncionesInyect.Estado = estado.MainContainer
	globals.FuncionesInyect.Transferencias = transferencias.MainContainer
	globals.FuncionesInyect.Rifas = rifas.MainContainer
	globals.FuncionesInyect.Anotaciones = anotaciones.MainContainer
	globals.FuncionesInyect.VerSocios = ver_usuarios.MainContainer
	globals.FuncionesInyect.Registros = registros.MainContainer
	globals.FuncionesInyect.ModificarSocios = modificar_usuarios.MainContainer
	globals.FuncionesInyect.Ajustes = ajustes.MainContainer
}
