package main

import (
	"fmt"
)

func main() {

	//Las constantes se guardan en memoria, se asignan una sola vez y 
	//seran datos de lectura. No se pueden modificar

	const myIntConst int = 12
	fmt.Println("Mi constante es: ", myIntConst)

	const myFirstStringConst = "a12"
	fmt.Println("Mi constante es: ", myFirstStringConst)

	/*
	myFirstStringConst = "test"
	fmt.Println("Mi constante es: ", myFirstStringConst)
	
	Da error por que no podemos modificar la constante
	*/

	const myStringConst string = "test"
	fmt.Println("Mi constante es: ", myStringConst)

	/*
	myStringConst2 string = 50
	fmt.Println("Mi constante es: ", myStringConst2)
	
	Da error por que le estamos pasando un int a un string
	*/

}