package main

import (
	// importaciones de la biblioteca standart
	"log"

	// importaciones de fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	// importaciones de mis paquetes
	"fondo/globals"
	myfn "fondo/misFunciones"
)

func make_slide_bar(win fyne.Window) *fyne.Container {

	// Paginas generales
	btnMenu := widget.NewButton(
		"📘 menu", func() {
			globals.PaginaActual = "menu"
			globals.Refresh()
		},
	)
	btnCuotas := widget.NewButton(
		"📆 cuotas", func() {
			globals.PaginaActual = "cuotas"
			globals.Refresh()
		},
	)
	btnPrestamos := widget.NewButton(
		"💵 prestamos", func() {
			globals.PaginaActual = "prestamos"
			globals.Refresh()
		},
	)
	btnEstado := widget.NewButton(
		"📈 estado", func() {
			globals.PaginaActual = "estado"
			globals.Refresh()
		},
	)
	btnTransferencias := widget.NewButton(
		"🏛 transferencias", func() {
			globals.PaginaActual = "transferencias"
			globals.Refresh()
		},
	)
	btnRifas := widget.NewButton(
		"💰 rifas", func() {
			globals.PaginaActual = "rifas"
			globals.Refresh()
		},
	)
	btnAnotaciones := widget.NewButton(
		"📝 anotaciones", func() {
			globals.PaginaActual = "anotaciones"
			globals.Refresh()
		},
	)
	btnVerSocios := widget.NewButton(
		"🔍 ver socios", func() {
			globals.PaginaActual = "ver usuarios"
			globals.Refresh()
		},
	)
	btnRegistros := widget.NewButton(
		"📚 registros", func() {
			globals.PaginaActual = "registros"
			globals.Refresh()
		},
	)

	btnModificarSocios := widget.NewButtonWithIcon(
		"modificar socios",
		theme.WarningIcon(),
		func() {
			globals.PaginaActual = "modificar usuarios"
			globals.Refresh()
		},
	)
	btnAjustes := widget.NewButtonWithIcon(
		"ajustes",
		theme.SettingsIcon(),
		func() {
			globals.PaginaActual = "ajustes"
			globals.Refresh()
		},
	)

	var adminContainer *fyne.Container
	var btnIngresar *widget.Button

	btnIngresar = widget.NewButtonWithIcon(
		"ingresar",
		theme.LoginIcon(),
		func() {
			password := widget.NewPasswordEntry()
			items := []*widget.FormItem{
				widget.NewFormItem("Password", password),
			}

			dialog.ShowForm(
				"Acceder como admimistrador",
				"Acceder",
				"Cancelar",
				items,
				func(b bool) {

					if password.Text == "1234" {
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
									globals.Admin = false
									log.Println("Modo administrador desactivado")
								},
							),
						}
						adminContainer.Refresh()
						globals.Admin = true
						log.Println("Modo administrador activado")
					}
				},
				win,
			)
		},
	)

	adminContainer = container.NewVBox(
		btnIngresar,
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
		entradaUser,
		botonBuscar,
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

		widget.NewLabelWithStyle(
			"Buscar Usuarios",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		setUser,
	)
}
