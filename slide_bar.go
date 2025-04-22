package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"fondo/paginas/cuotas"
)

func make_slide_bar(cont *fyne.Container) *fyne.Container {
	// Paginas generales
	btnMenu := widget.NewButton(
		"menu",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnCuotas := widget.NewButton(
		"cuotas",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnPrestamos := widget.NewButton(
		"prestamos",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnEstado := widget.NewButton(
		"estado",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnTransferencias := widget.NewButton(
		"transferencias",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnRifas := widget.NewButton(
		"rifas",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnAnotaciones := widget.NewButton(
		"anortaciones",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnVerSocios := widget.NewButton(
		"ver socios",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnRegistros := widget.NewButton(
		"registros",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)

	// Paginas administrativas

	btnModificarSocios := widget.NewButton(
		"modificar socios",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)
	btnAjustes := widget.NewButton(
		"ajustes",
		func() {
			cont.Objects = []fyne.CanvasObject{} // Limpiar el contenido existente
			cont.Add(cuotas.MainContainer())     // Agregar nuevo contenido
			cont.Refresh()                       // Actualizar el contenedor para mostrar el nuevo contenido
		},
	)

	return container.NewVBox(
		widget.NewLabelWithStyle(
			"Paginas Generales",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		btnMenu,
		btnCuotas,
		btnPrestamos,
		btnEstado,
		btnTransferencias,
		btnRifas,
		btnAnotaciones,
		btnVerSocios,
		btnRegistros,
		container.NewVBox(
			widget.NewLabelWithStyle(
				"Paginas Administrativas",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			btnModificarSocios,
			btnAjustes,
		),
	)
}
