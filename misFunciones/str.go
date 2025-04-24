package misfunciones

import (
	"strings"
)

func Title(texto string) string {
	textoR := []rune(texto)

	makeUpper := true

	for i, c := range textoR {
		if c == ' ' {
			makeUpper = true
			continue
		}
		if makeUpper {
			textoR[i] = []rune(strings.ToUpper(string(c)))[0]
			makeUpper = false
		}
	}
	return string(textoR)
}
