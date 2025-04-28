package misfunciones

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MakeTable(data [][]string) *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)
}

func MakeTableCuotas(data [51][4]string) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return 51, 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)
	table.SetColumnWidth(0, 70)
	table.SetColumnWidth(1, 110)
	table.SetColumnWidth(2, 70)
	table.SetColumnWidth(3, 70)
	return table
}

func MakeTable4(data [][4]string, ajustes [4]float32) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 4 // filas, columnas
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	table.SetColumnWidth(0, ajustes[0])
	table.SetColumnWidth(1, ajustes[1])
	table.SetColumnWidth(2, ajustes[2])
	table.SetColumnWidth(3, ajustes[3])
	return table
}
