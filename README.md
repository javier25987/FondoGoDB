# Proyecto fondo san javier

> Link al repositorio [padre](https://github.com/javier25987/FondoGoDB)

## Anotaciones

> Para cambiar la escala de la ventana lo recomendable cambiar la escala de al pantalla `$env:FYNE_SCALE = "1.2"`

> En el fondo hay una gerarquia de simbolos para separar los elementos de un string la gerarquia es la siguiente `_ >> # >> ?`

> El simbolo `/` se reserva unicamente para las fechas y para nada mas

## Deuda tecnica

* algoritmo para identificar boletas repetidas (seccion de rifas)

## Secciones finalizadas

* [ ]  arranque
* [ ]  menu
* [ ]  cuotas
* [ ]  prestamos
* [ ]  rifas
* [ ]  registros
* [ ]  transferencias
* [ ]  ver socios
* [ ]  Anotaciones
* [ ]  ajustes (temporal)
* [ ]  modificar socios
* [ ]  analizar usuarios

## Errores de el programa

Estoy reaciendo el programa en go esta seccion esta vacia por el momento

## Codigo que elimine y podria servir en un futuro

```go
disabled := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(string) {})
disabled.Horizontal = true
disabled.Disable()
largeText := widget.NewMultiLineEntry()
password := widget.NewPasswordEntry()
password.SetPlaceHolder("Password")
```


para incluir datos en un formato (como el de pago en cuotas) se usa `form.Append("<Text>", <Widget>)`

## Datos a tener en cuenta para mostara en la tabla

* cuanto ha pagado
* cuotas que tiene pagas
* cunto ha pagado en multas
* multas adeudas
* prestamos activos
* dinero pagado en intereses
* deuda de prestamos activos

## Agradecimiento

Este proyecto fue hecho para mi padre al cual le agradezo todo lo que me ha dado y la educacion que me esta pagando ya que gracias a eso obtuve los conocimientos para realizar este proyecto,

GRACIAS PAPA

> **NOTA:** no es necesario que el programa calcule todo asincronicamente, lo puede hacer el siguiente dia al iniciarse