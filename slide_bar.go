package main

import (
	// importaciones de la biblioteca standart
	"log"

	// importaciones de fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	// importaciones de mis paquetes
	"fondo/global"
)

func make_slide_bar(win fyne.Window) *widget.Tree {

	// Estado dinámico de las ramas
	childrenMap := map[string][]string{
		"": {"Paginas generales:", "Paginas administrativas:"},
		"Paginas generales:": {
			// "📘 menu",
			"📆 cuotas",
			"💵 prestamos",
			"📈 estado",
			"🏛 transferencias",
			"💰 rifas",
			"📝 anotaciones",
			"🔍 ver socios",
			"📚 registros",
		},
		"Paginas administrativas:": {"Ingresar"},
	}

	// Crear el árbol
	tree := widget.NewTree(
		// Función para obtener los hijos de un nodo
		func(uid string) []string {
			return childrenMap[uid]
		},
		// Función para verificar si es un nodo o una hoja
		func(uid string) bool {
			_, ok := childrenMap[uid]
			return ok
		},
		// Crear el nodo
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Nodo") // Crear una etiqueta para los nodos
		},
		// Actualizar el nodo
		func(uid string, branch bool, obj fyne.CanvasObject) {
			// Aquí utilizamos `uid` para actualizar el nodo
			obj.(*widget.Label).SetText(uid) // Actualizamos el texto con el `uid`
		},
	)

	tree.OpenAllBranches()

	// Acción al seleccionar un nodo
	tree.OnSelected = func(uid string) {

		refreshCont := true

		switch uid {
		case "Ingresar":

			password := widget.NewPasswordEntry()
			items := []*widget.FormItem{
				widget.NewFormItem("Clave:", password),
			}

			dialog.ShowForm(
				"Acceder como admimistrador",
				"Acceder",
				"Cancelar",
				items,
				func(b bool) {
					if password.Text == "1234" {
						childrenMap["Paginas administrativas:"] = []string{
							"✏️ modificar usuarios", "⚙️ ajustes", "Salir",
						}
						global.Admin = true
						log.Println("Modo administrador activado")
					}
				},
				win,
			)
			refreshCont = false

		case "Salir":
			childrenMap["Paginas administrativas:"] = []string{"Ingresar"}
			log.Println("Modo administrador desactivado")
			refreshCont = false

		case "📘 menu":
			global.PaginaActual = "menu"
		case "📆 cuotas":
			global.PaginaActual = "cuotas"
		case "💵 prestamos":
			global.PaginaActual = "prestamos"
		case "📈 estado":
			global.PaginaActual = "estado"
		case "🏛 transferencias":
			global.PaginaActual = "transferencias"
		case "💰 rifas":
			global.PaginaActual = "rifas"
		case "📝 anotaciones":
			global.PaginaActual = "anotaciones"
		case "🔍 ver socios":
			global.PaginaActual = "ver usuarios"
		case "📚 registros":
			global.PaginaActual = "registros"
		case "✏️ modificar usuarios":
			global.PaginaActual = "modificar usuarios"
		case "⚙️ ajustes":
			global.PaginaActual = "ajustes"
		}

		if refreshCont {
			global.Refresh()
		} else {
			tree.Refresh()
		}
	}

	return tree
}
