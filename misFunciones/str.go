package misfunciones

import (
	"errors"
	"fmt"
	"fondo/globals"
	"strconv"

	"fyne.io/fyne/v2/dialog"
)

func FormatComas(n int) string {
	numberStr := strconv.Itoa(n)

	if len(numberStr) < 4 {
		return numberStr
	}

	// Convertir el string a runas para manejar correctamente los caracteres
	runes := []rune(numberStr)
	length := len(runes)

	// Calcular cuántas comas necesitaremos
	commaCount := (length - 1) / 3

	// Crear un nuevo slice con capacidad suficiente
	result := make([]rune, length+commaCount)

	// Índices para recorrer el string original y el resultado
	i, j := length-1, len(result)-1
	count := 0

	for i >= 0 {
		if count == 3 {
			result[j] = ','
			j--
			count = 0
		}
		result[j] = runes[i]
		j--
		i--
		count++
	}

	return string(result)
}

func RectNumber(num string) (bool, int) {
	numero, err := strconv.Atoi(num)

	if err != nil {
		mensaje := fmt.Sprintf(
			"[%s] No es un valor válido", num,
		)
		dialog.ShowError(errors.New(mensaje), globals.WinDialog)

		return false, 0
	}

	return true, numero
}
