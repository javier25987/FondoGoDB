/*
La simple existencia de este paquete es una aberracion ya que es una mala prectica
pero para mi me sirve y es funcional

este paquete agrega una capa mas de complejidad al entorno ya que puede crear
condiciones de carrera por esto la convenciones que solo el paquete main y el
paquete inyect pueden cargar datos aca y los demas paquetes solo pueden usar
la funcion `Refresh()`
*/
package globals

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Funciones struct {
	Menu            func() *container.Split
	Cuotas          func() *container.Split
	Prestamos       func() *container.Split
	Estado          func() *container.Split
	Transferencias  func() *container.Split
	Rifas           func() *container.Split
	Anotaciones     func() *container.Split
	VerSocios       func() *container.Split
	Registros       func() *container.Split
	ModificarSocios func() *container.Split
	Ajustes         func() *container.Split
}

var Admin bool = false // Variable para determinar si es administrador o no
var Index int = 0      // determina el actual del fondo
var PaginaActual string = ""
var MyWindow *fyne.Window
var Container2 *fyne.Container

var FuncionesInyect Funciones

func Refresh() {

	var NewContainer *container.Split

	switch PaginaActual {

	case "cuotas":
		NewContainer = FuncionesInyect.Cuotas()
	case "prestamos":
		NewContainer = FuncionesInyect.Prestamos()
	case "estado":
		NewContainer = FuncionesInyect.Estado()
	case "transferencias":
		NewContainer = FuncionesInyect.Transferencias()
	case "rifas":
		NewContainer = FuncionesInyect.Rifas()
	case "anotaciones":
		NewContainer = FuncionesInyect.Anotaciones()
	case "ver_usuarios":
		NewContainer = FuncionesInyect.VerSocios()
	case "registros":
		NewContainer = FuncionesInyect.Registros()
	case "modificar_usuarios":
		NewContainer = FuncionesInyect.ModificarSocios()
	case "ajustes":
		NewContainer = FuncionesInyect.Ajustes()
	case "menu":
		NewContainer = FuncionesInyect.Menu()
	}

	Container2.Objects = []fyne.CanvasObject{}
	Container2.Add(NewContainer)
	Container2.Refresh()
}
