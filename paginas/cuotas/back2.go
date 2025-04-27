package cuotas

import (
	"fmt"
	"strconv"
	"strings"

	mySQL "fondo/sql"
)

func pagarCuotas(index int, cuotas int, multas int, metodo string) {
	fmt.Printf(
		"El usuario %d ha pagado %d semanas y %d: metodo de pago %s\n",
		index, cuotas, multas, metodo,
	)
}

func rectificarPago(cuotas int, multas int, metodo string) (bool, string) {
	if metodo == "" {
		return false, "Selecione un metodo de pago"
	}
	if cuotas+multas <= 0 {
		return false, "No se entiende que desea pagar"
	}
	return true, ""
}

func rectificarBloqueo(rect string) (bool, string) {
	numRect, err := strconv.Atoi(rect)

	if err != nil {
		return false, fmt.Sprintf("[%s] no se puede convertir a un numero", rect)
	}

	if numRect < 1 || numRect > 50 {
		return false, "El valor ingresado no esta en el intervalo estipulado"
	}

	return true, ""
}

/*
Esta funcion recibe un string ya que todo el proceso anterior de certificacion se
asegura de solo recibir una cifra entre 1 y 50 codificada como un string
*/
func makeBlock(index int, num string) {

	block := mySQL.GetValueStr("cuotas", "bloqueos", index)

	if block == "n" {
		// cargar 'num' a la base
		mySQL.SetValueStr("cuotas", "bloqueos", index, num)
		return
	}

	if num == block {
		// cargamos una 'n' a la base
		mySQL.SetValueStr("cuotas", "bloqueos", index, "n")
		return
	}

	blockArray := strings.Split(block, "_")
	app := true

	for i, j := range blockArray {
		if j == num {
			blockArray[0], blockArray[i] = blockArray[i], blockArray[0]
			app = false
			break
		}
	}

	if app {
		blockArray = append(blockArray, num)
	} else {
		blockArray = blockArray[1:]
	}

	exit := strings.Join(blockArray, "_")

	// cargamos exit a la base de datos
	mySQL.SetValueStr("cuotas", "bloqueos", index, exit)
}
