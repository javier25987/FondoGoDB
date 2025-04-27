package inyect

import (
	"fondo/globals"

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

func Inyect() {
	globals.FuncionesInyect.Cuotas = cuotas.MainContainer
	globals.FuncionesInyect.Prestamos = prestamos.MainContainer
	globals.FuncionesInyect.Estado = estado.MainContainer
	globals.FuncionesInyect.Transferencias = transferencias.MainContainer
	globals.FuncionesInyect.Rifas = rifas.MainContainer
	globals.FuncionesInyect.Anotaciones = anotaciones.MainContainer
	globals.FuncionesInyect.VerSocios = ver_usuarios.MainContainer
	globals.FuncionesInyect.Registros = registros.MainContainer
	globals.FuncionesInyect.ModificarSocios = modificar_usuarios.MainContainer
	globals.FuncionesInyect.Ajustes = ajustes.MainContainer
}
