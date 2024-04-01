package main

import (
	"fmt"
	"unsafe"
)

func main() {

	/* ******************* 	Arrays ******************* */

	fmt.Println(`
	******************* Arrays *******************
	`)

	//Arrays --> secuencia de valores. Son secuencias consecutivas. Todos los valores deben tener el mismo tipo de dato (int - str)
	//Estan representados por índices, es decir, cada posición del array. Siempre se arrancará en el índice 0
	//Para acceder a un valor del array: vectorName[index]
	//Se puede usar cualquier tipo de dato dentro del array
	//No se puede modificar el tamaño del array configurado [5]int

	var myIntVar int = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d  \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar1 [5]int //array compuesto por valores integer. Son 5 valores dentro del array (n - 1)

	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1)) // [0 0 0 0 0] - size: 5

	myArrayVar2 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar2, " - size: ", len(myArrayVar2))

	myArrayVar1[0] = 2 //En índice 0 le pongo valor 0
	myArrayVar1[1] = 5
	myArrayVar1[2] = 9
	//myArrayVar1[7] = 12 --> da error, por que es un índice no válido en el array

	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1), " - get index 1: ", myArrayVar1[1])
	fmt.Println("En el índice 0 del array se encuentra el valor: ", myArrayVar1[0])

	fmt.Printf("type: %T, value: %v, bytes: %d, bits: %d  \n", myArrayVar1, myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	fmt.Println(`
	******************* Recorrer arrays con for *******************
	`)

	for i := range myArrayVar1 { //Iria de 0 a 4 (n-1)
		fmt.Println("index: ", i, " - value: ", myArrayVar1[i])
	}

	for i, v := range myArrayVar1 { // i indice, v valor
		fmt.Println("index: ", i, " - value: ", v)
	}

	for _, v := range myArrayVar1 { // _ omitimos el indice por que no lo usamos
		fmt.Println("value: ", v)
	}

}
