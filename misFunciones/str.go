package misfunciones

import "strconv"

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
