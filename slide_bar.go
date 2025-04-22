package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	btnMenu := widget.NewButtonWithIcon(
		"menu",
		theme.HomeIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(cuotas.MainContainer())
			cont.Refresh()
		},
	)
	btnCuotas := widget.NewButtonWithIcon(
		"cuotas",
		theme.CalendarIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(cuotas.MainContainer())
			cont.Refresh()
		},
	)
	btnPrestamos := widget.NewButtonWithIcon(
		"prestamos",
		theme.CalendarIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(prestamos.MainContainer())
			cont.Refresh()
		},
	)
	btnEstado := widget.NewButtonWithIcon(
		"estado",
		theme.AccountIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(estado.MainContainer())
			cont.Refresh()
		},
	)
	btnTransferencias := widget.NewButtonWithIcon(
		"transferencias",
		theme.ListIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(transferencias.MainContainer())
			cont.Refresh()
		},
	)
	btnRifas := widget.NewButtonWithIcon(
		"rifas",
		theme.CalendarIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(rifas.MainContainer())
			cont.Refresh()
		},
	)
	btnAnotaciones := widget.NewButtonWithIcon(
		"anotaciones",
		theme.DocumentCreateIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(anotaciones.MainContainer())
			cont.Refresh()
		},
	)
	btnVerSocios := widget.NewButtonWithIcon(
		"ver socios",
		theme.SearchIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(ver_usuarios.MainContainer())
			cont.Refresh()
		},
	)
	btnRegistros := widget.NewButtonWithIcon(
		"registros",
		theme.StorageIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(registros.MainContainer())
			cont.Refresh()
		},
	)
	btnModificarSocios := widget.NewButtonWithIcon(
		"modificar socios",
		theme.WarningIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(modificar_usuarios.MainContainer())
			cont.Refresh()
		},
	)
	btnAjustes := widget.NewButtonWithIcon(
		"ajustes",
		theme.SettingsIcon(),
		func() {
			cont.Objects = []fyne.CanvasObject{}
			cont.Add(ajustes.MainContainer())
			cont.Refresh()
		},
	)

	var adminContainer *fyne.Container
	var btnIngresar *widget.Button

	btnIngresar = widget.NewButtonWithIcon(
		"ingresar",
		theme.LoginIcon(),
		func() {
			adminContainer.Objects = []fyne.CanvasObject{
				btnModificarSocios,
				btnAjustes,
				widget.NewButtonWithIcon(
					"salir",
					theme.LogoutIcon(),
					func() {
						adminContainer.Objects = []fyne.CanvasObject{}
						adminContainer.Add(btnIngresar)
						adminContainer.Refresh()
						Admin = false
					},
				),
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
