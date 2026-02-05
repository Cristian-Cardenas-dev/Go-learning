package main

import "fmt"

func main() {
	// Forma explicita con tipo
	var nombre string = "Juan"
	var edad int = 25
	fmt.Println("Nombre:", nombre, "Edad:", edad)

	// Go infiere el tipo y solo puede usarse dentro de funciones
	ciudad := "Madrid"
	temperatura := 22.5
	fmt.Println("Ciudad:", ciudad, "Temp:", temperatura)

	// int, float64, bool, string
	var numero int = 42
	var decimal float64 = 3.14
	var activo bool = true
	fmt.Println(numero, decimal, activo)

	// Variables sin valor inicial tienen valores por defecto, el 0 de cada valor
	var sinValor int      // 0
	var textoVacio string // ""
	var sinBooleano bool  // false
	fmt.Println("Zero values:", sinValor, textoVacio, sinBooleano)

	// Go NO convierte automaticamente, entre valores de tipo nunmericos
	var entero int = 42
	var flotante float64 = float64(entero)
	fmt.Println("Conversion:", entero, "->", flotante)
}
