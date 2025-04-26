package cuotas

import (
	"fmt"
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
