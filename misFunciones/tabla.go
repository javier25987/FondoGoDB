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
