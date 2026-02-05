package main

import "fmt"

func main() {
	// Evalua condicion y ejecuta codigo
	edad := 20
	if edad >= 18 {
		fmt.Println("Mayor de edad")
	} else {
		fmt.Println("Menor de edad")
	}

	// Seleccion multiple sin break
	dia := "Lunes"
	switch dia {
	case "Lunes":
		fmt.Println("Inicio de semana")
	case "Sabado", "Domingo":
		fmt.Println("Fin de semana")
	default:
		fmt.Println("Dia normal")
	}

	// Unico tipo de loop en Go
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Salta a la siguiente iteracion
	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			continue // salta pares
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Itera sobre colecciones
	colores := []string{"Rojo", "Verde", "Azul"}
	for i, color := range colores {
		fmt.Println(i, "-", color)
	}
}
