package main

import (
	"composicion_punteros_interfaces/pointers"
	"fmt"
)

func main() {
	fmt.Println("Main")
	/* ******************* 	Composición, punteros e interfaces ******************* */
	fmt.Println(`
	******************* Ejemplo sin y con punteros *******************
	`)

	// Declarar una variable
	myVar := 30

	fmt.Println("Dirección de memoria de var fuera de la función increment: ", &myVar)
	// Llamar a la función increment y pasarle la variable por valor

	fmt.Println(`
	******************* Punteros *******************
	`)

	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println("Dirección memoría variable i: ", &i)
	fmt.Println("Puntero: ", iP)
	fmt.Println("Dirección memoría del puntero: ", &iP)
	fmt.Println("Variable i: ", i)
	fmt.Println("Valor del puntero: ", *iP)

	//Modifico el valor de iP
	*iP = 1

	//Cuando modifico uno se modifica el otro
	fmt.Printf("value i: %d, value pointer i: %d\n", i, *iP)

	fmt.Println()

	myVar = 30
	fmt.Printf("address: %p, value: %d\n", &myVar, myVar)

	fmt.Println()

	pointers.Increment(myVar)

	fmt.Println("Dirección de memoria de var fuera de la función increment: ", &myVar)

	fmt.Println()

	pointers.IncrementP(&myVar) // Llama a la función IncrementP pasando un puntero a myVar

	fmt.Println("Nuevo valor de myVar fuera de IncrementP:", myVar)
}
