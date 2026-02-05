package main

import (
	"errors"
	"fmt"
)

// COMPORTAMIENTO: Ejecuta acciones pero NO devuelve ningun valor
// CASOS DE USO:
//   - Imprimir mensajes o logs
//   - Modificar variables globales
//   - Efectos secundarios (escribir archivos, enviar emails, etc)
//   - Cuando solo necesitas que "haga algo" sin esperar resultado
func saludar() {
	fmt.Println("Hola!")
}

// COMPORTAMIENTO: Recibe datos, hace algo con ellos, pero NO devuelve nada
// CASOS DE USO:
//   - Mostrar informacion formateada
//   - Validar y mostrar errores
//   - Guardar datos (en BD, archivo)
//   - Enviar notificaciones
func saludarPersona(nombre string) {
	fmt.Println("Hola", nombre)
}

// FUNCION CON RETORNO
func sumar(a, b int) int {
	return a + b
}

// MULTIPLES RETORNOS
// Patron comun: (valor, error)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division por cero")
	}
	return a / b, nil
}

// FUNCION DEFER
// Se ejecuta al final de la funcion
func ejemploDefer() {
	defer fmt.Println("Print Ultimo")
	fmt.Println("Print Primero")
}

func main() {
	// Funciones sin retorno - solo ejecutan acciones
	saludar()
	saludarPersona("Carlos")

	// Funcion con retorno
	resultado := sumar(5, 3)
	fmt.Println("Suma:", resultado)

	// Multiples retornos - manejar error
	cociente, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", cociente)
	}

	// Defer
	ejemploDefer()

	// FUNCION ANONIMA
	// Funcion sin nombre
	multiplicar := func(a, b int) int {
		return a * b
	}
	fmt.Println("Multiplicar:", multiplicar(4, 5))
}
