package main

import (
	"fmt"
)

func main() {

	//En las variable almacenamos valores que podemos modificar y consultar

	var myIntVar int // Por defecto valor de variable = 0

	myIntVar = -12

	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint // Solo valores positivos

	myUintVar = 12

	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string // Por defecto vacío

	myStringVar = "my string variable"

	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool // Por defecto false

	myBoolVar = true

	fmt.Println("My variable is: ", myBoolVar)


	//var myOtherVal string // No compila por que no se esta usando

	//GO nos obliga a utilizar todas las variables que intanciamos. 
	//GO no compila si encuentra variables que no se están utilizando en 
	//ninguna parte del código

	//Dirección de memoria de la variable
	//a la valriable se le debe agregar &___________
	fmt.Println("My variable address is: ", &myStringVar)

	myIntVar2 := 12 // := nos permite instanciar y agregar valor a la variable
					// toma el dato de la variable según lo que se le pase
					// En ese caso toma que es un int

	fmt.Println("My variable is: ", myIntVar2)

	myStringVar2 := "My string variable with :="
	fmt.Println("My variable is: ", myStringVar2)

	myBoolVar2 := true
	fmt.Println("My variable is: ", myBoolVar2)

}