package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fondo/globals"
	"fondo/inyect"
)

// "fyne.io/fyne/v2/layout"

func main() {

	log.Println("Iniciando el programa...")

	myApp := app.New()
	myWindow := myApp.NewWindow("Sidebar Demo")
	globals.MyWindow = &myWindow

	inyect.Inyect()

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
