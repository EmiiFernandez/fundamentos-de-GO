package main

import (
	"composicion_punteros_interfaces/pointers"
	"fmt"
)

func main() {
	fmt.Println("Main")
	/* ******************* 	Composición, punteros e interfaces ******************* */

	fmt.Println(`
	******************* Punteros *******************
	`)

	// Declarar una variable
	myVar := 30

	fmt.Println("Dirección de memoria de var fuera de la función increment: ", &myVar)
	// Llamar a la función increment y pasarle la variable por valor

	pointers.Increment(myVar)

	fmt.Println("Dirección de memoria de var fuera de la función increment: ", &myVar)

	pointers.IncrementP(&myVar) // Llama a la función IncrementP pasando un puntero a myVar

	fmt.Println("Nuevo valor de myVar fuera de IncrementP:", myVar)
}
