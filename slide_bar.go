package main

import (
	// importaciones de la biblioteca standart
	"errors"
	"fmt"
	"log"
	"strconv"

	// importaciones de fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	// importaciones de mis paquetes
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

func make_slide_bar(cont *fyne.Container, win fyne.Window) *fyne.Container {

	// funciones de paginas
	funMenu := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(widget.NewLabel("Menu"))
		cont.Refresh()
		PaginaActual = "menu"
	}
	funCuotas := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(cuotas.MainContainer(Index))
		cont.Refresh()
		PaginaActual = "cuotas"
	}
	funPrestamos := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(prestamos.MainContainer())
		cont.Refresh()
		PaginaActual = "prestamos"
	}
	funEstado := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(estado.MainContainer())
		cont.Refresh()
		PaginaActual = "estado"
	}
	funTransferencias := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(transferencias.MainContainer())
		cont.Refresh()
		PaginaActual = "transferencias"
	}
	funRifas := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(rifas.MainContainer())
		cont.Refresh()
		PaginaActual = "rifas"
	}
	funAnotaciones := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(anotaciones.MainContainer())
		cont.Refresh()
		PaginaActual = "anotaciones"
	}
	funVerUsuarios := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(ver_usuarios.MainContainer())
		cont.Refresh()
		PaginaActual = "ver_usuarios"
	}
	funRegistros := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(registros.MainContainer())
		cont.Refresh()
		PaginaActual = "registros"
	}
	funModificarUsuarios := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(modificar_usuarios.MainContainer())
		cont.Refresh()
		PaginaActual = "modificar_usuarios"
	}
	funAjustes := func() {
		cont.Objects = []fyne.CanvasObject{}
		cont.Add(ajustes.MainContainer())
		cont.Refresh()
		PaginaActual = "ajustes"
	}

	// Paginas generales
	btnMenu := widget.NewButton("📘 menu", funMenu)
	btnCuotas := widget.NewButton("📆 cuotas", funCuotas)
	btnPrestamos := widget.NewButton("💵 prestamos", funPrestamos)
	btnEstado := widget.NewButton("📈 estado", funEstado)
	btnTransferencias := widget.NewButton("🏛 transferencias", funTransferencias)
	btnRifas := widget.NewButton("💰 rifas", funRifas)
	btnAnotaciones := widget.NewButton("📝 anotaciones", funAnotaciones)
	btnVerSocios := widget.NewButton("🔍 ver socios", funVerUsuarios)
	btnRegistros := widget.NewButton("📚 registros", funRegistros)

	btnModificarSocios := widget.NewButtonWithIcon(
		"modificar socios",
		theme.WarningIcon(),
		funModificarUsuarios,
	)
	btnAjustes := widget.NewButtonWithIcon(
		"ajustes",
		theme.SettingsIcon(),
		funAjustes,
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
									Admin = false
									log.Println("Modo administrador desactivado")
								},
							),
						}
						adminContainer.Refresh()
						Admin = true
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

	botonBuscar := widget.NewButton("🔎 Buscar", func() {
		numeroUser, err := strconv.Atoi(entradaUser.Text)

		if err != nil {
			mensaje := fmt.Sprintf(
				"[%s] No es un valor válido",
				entradaUser.Text,
			)
			dialog.ShowError(errors.New(mensaje), win)
		}

		Index = numeroUser

		switch PaginaActual {

		case "cuotas":
			funCuotas()
		case "prestamos":
			funPrestamos()
		case "estado":
			funEstado()
		case "transferencias":
			funTransferencias()
		case "rifas":
			funRifas()
		case "anotaciones":
			funAnotaciones()
		case "ver_usuarios":
			funVerUsuarios()
		case "registros":
			funRegistros()
		case "modificar_usuarios":
			funModificarUsuarios()
		case "ajustes":
			funAjustes()

		default:
			funMenu()
		}
	})

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
