package main

import (
	// importaciones de la biblioteca standart
	"log"

	// importaciones de fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	// importaciones de mis paquetes
	"fondo/global"
	"fondo/pkg/ajustes"
	"fondo/pkg/anotaciones"
	"fondo/pkg/cuotas"
	"fondo/pkg/estado"
	"fondo/pkg/menu"
	"fondo/pkg/modificar_usuarios"
	"fondo/pkg/prestamos"
	"fondo/pkg/registros"
	"fondo/pkg/rifas"
	"fondo/pkg/transferencias"
	"fondo/pkg/ver_usuarios"
)

func main() {
	log.Println("Iniciando el programa...")

	myApp := app.New()
	myWindow := myApp.NewWindow("Fondo San Javier")
	global.MyWindow = &myWindow
	global.WinDialog = myWindow

	Inyect()

	// Panel principal que irá cambiando con la navegación
	myContainer := container.NewPadded()
	global.Container2 = myContainer

	// Contenedor general: Sidebar + Contenido principal
	mainContainer := container.NewHSplit(
		make_slide_bar(myWindow), myContainer,
	)
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
	global.FuncionesInyect.Menu = menu.MainContainer
	global.FuncionesInyect.Cuotas = cuotas.MainContainer
	global.FuncionesInyect.Prestamos = prestamos.MainContainer
	global.FuncionesInyect.Estado = estado.MainContainer
	global.FuncionesInyect.Transferencias = transferencias.MainContainer
	global.FuncionesInyect.Rifas = rifas.MainContainer
	global.FuncionesInyect.Anotaciones = anotaciones.MainContainer
	global.FuncionesInyect.VerSocios = ver_usuarios.MainContainer
	global.FuncionesInyect.Registros = registros.MainContainer
	global.FuncionesInyect.ModificarSocios = modificar_usuarios.MainContainer
	global.FuncionesInyect.Ajustes = ajustes.MainContainer
}
