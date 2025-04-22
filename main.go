package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// "fyne.io/fyne/v2/layout"

var Admin bool = false // Variable para determinar si es administrador o no

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Sidebar Demo")

	// Panel principal que irá cambiando con la navegación
	myContainer := container.NewPadded()

	// Sidebar
	sidebar := make_slide_bar(myContainer)

	// Contenedor general: Sidebar + Contenido principal
	mainContainer := container.NewHSplit(sidebar, myContainer)
	mainContainer.SetOffset(0.1) // Tamaño relativo de la barra lateral

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}
