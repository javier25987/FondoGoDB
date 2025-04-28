package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

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
)

// "fyne.io/fyne/v2/layout"

func main() {

	log.Println("Iniciando el programa...")

	myApp := app.New()
	myWindow := myApp.NewWindow("Fondo San Javier")
	globals.MyWindow = &myWindow

	Inyect()

	// Panel principal que irá cambiando con la navegación
	myContainer := container.NewPadded()
	globals.Container2 = myContainer

	// Sidebar
	sidebar := container.NewVScroll(
		make_slide_bar(myContainer, myWindow),
	)

	// Contenedor general: Sidebar + Contenido principal
	mainContainer := container.NewHSplit(sidebar, myContainer)
	mainContainer.SetOffset(0.1) // Tamaño relativo de la barra lateral

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

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
