package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"fondo/paginas/ajustes"
	"fondo/paginas/anotaciones"
	"fondo/paginas/cuotas"
	"fondo/paginas/estado"
	"fondo/paginas/modificar_usuarios"
	"fondo/paginas/prestamos"
	"fondo/paginas/registros"
	"fondo/paginas/rifas"
	"fondo/paginas/transferencias"
	"fondo/paginas/ver_usuarios"
)

func make_slide_bar(cont *fyne.Container) *fyne.Container {
	// Paginas generales
	btnMenu := widget.NewButton(
		"menu",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(cuotas.MainContainer())
			cont.Refresh()
		},
	)
	btnCuotas := widget.NewButton(
		"cuotas",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(cuotas.MainContainer())
			cont.Refresh()
		},
	)
	btnPrestamos := widget.NewButton(
		"prestamos",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(prestamos.MainContainer())
			cont.Refresh()
		},
	)
	btnEstado := widget.NewButton(
		"estado",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(estado.MainContainer())
			cont.Refresh()
		},
	)
	btnTransferencias := widget.NewButton(
		"transferencias",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(transferencias.MainContainer())
			cont.Refresh()
		},
	)
	btnRifas := widget.NewButton(
		"rifas",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(rifas.MainContainer())
			cont.Refresh()
		},
	)
	btnAnotaciones := widget.NewButton(
		"anortaciones",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(anotaciones.MainContainer())
			cont.Refresh()
		},
	)
	btnVerSocios := widget.NewButton(
		"ver socios",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(ver_usuarios.MainContainer())
			cont.Refresh()
		},
	)
	btnRegistros := widget.NewButton(
		"registros",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(registros.MainContainer())
			cont.Refresh()
		},
	)
	btnModificarSocios := widget.NewButton(
		"modificar socios",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(modificar_usuarios.MainContainer())
			cont.Refresh()
		},
	)
	btnAjustes := widget.NewButton(
		"ajustes",
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(ajustes.MainContainer())
			cont.Refresh()
		},
	)

	var adminContainer *fyne.Container
	var btnIngresar *widget.Button

	btnIngresar = widget.NewButton(
		"ingresar", func() {
			adminContainer.Objects = []fyne.CanvasObject{
				btnModificarSocios,
				btnAjustes,
				widget.NewButton("salir", func() {
					adminContainer.Objects = []fyne.CanvasObject{}
					adminContainer.Add(btnIngresar)
					adminContainer.Refresh()
					Admin = false
				}),
			}
			adminContainer.Refresh()
			Admin = true
		},
	)

	adminContainer = container.NewVBox(
		btnIngresar,
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
		widget.NewLabelWithStyle(
			"Paginas Administrativas",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		adminContainer,
	)
}
