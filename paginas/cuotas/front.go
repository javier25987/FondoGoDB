package cuotas

import (
	myfn "fondo/misFunciones"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContainer() *container.Split {
	data := [][]string{
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
		{"Nombre", "Edad", "Ciudad"},
		{"Ana", "25", "Bogotá"},
		{"Luis", "30", "Medellín"},
		{"Carlos", "28", "Cali"},
	}

	table := myfn.MakeTable(data)

	allContainer := container.NewHSplit(
		container.NewHScroll(table),
		container.NewVBox(
			widget.NewLabel("Contenido adicional"),
			widget.NewLabel("Más información"),
			widget.NewLabel("Detalles"),
		),
	)

	return allContainer
}
