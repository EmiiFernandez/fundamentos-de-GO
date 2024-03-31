package main

import (
	"fmt"
	"unsafe"
)

func main() {

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

	/* ******************* 	String Variable Size ******************* */

	//Son una secuencia de bytes

	var myStringVar string = "test1234"

	fmt.Printf("mi valor %s \n", myStringVar)
	//%s valor de tipo string

	myStringVar2 := `Mi variable de tipo texto en GO
		con múltiples
	líneas
=)`
	//Para poder crear texto con salto de linea, lo hago con ``

	fmt.Printf("mi valor %s \n", myStringVar2)

}
