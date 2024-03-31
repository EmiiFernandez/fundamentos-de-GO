package main

import (
	"fmt"
	"unsafe"
)

func main() {

	/* ******************* 	Variables ******************* */

	var myIntVar int // Por defecto valor de variable = 0

	/*
		Type		Size			Range
		int8		8 bits			-128 to 127
		int16		16 bits 		-2^15 to 2^16 -1
		int32		32 bits			-2^31 to 2^31 -1
		int64		64 bits 		-2^63 to 2^63 -1
		int	Platform dependente Patform dependent

	*/

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

	/* ******************* 	Unit Variable Size ******************* */

	var my8BitsUnitVar uint8 = 20
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n ", my8BitsUnitVar, my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var my16BitsUnitVar uint16 = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n ", my16BitsUnitVar, my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	var my32BitsUnitVar uint32 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n ", my32BitsUnitVar, my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	/*
		Type		Size			Range
		uint8		8 bits			0 to 255
		uint16		16 bits 		0 to 2^16 -1
		uint32		32 bits			0 to 2^32 -1
		uint64		64 bits 		0 to 2^64 -1
		uint	Platform dependente Patform dependent

		//uint --> aceptan valores solo de tipos positivos
	*/

	/* ******************* 	Float Variable Size ******************* */

	//Si o si tenemos que definir el espacio que ocupará el flotante
	//Solo hay de 32 y de 64 bits

	var myFloat32Var float32 = 3.14 //por defecto 0
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n ", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 590.12435
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n ", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	myOtherFloat := 7654.123
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n ", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat)*8)

}
