package main

import (
	// importaciones de la biblioteca standart
	"fmt"
	"log"

	// importaciones de fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

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
)

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

/*
Esta funcion lo unico que hace es inyectar las funciones principales a una
estructura en el paquete global
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

func Saludar() {
	fmt.Println("hola a todos")
}
