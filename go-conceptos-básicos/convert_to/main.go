package main

import (
	"fmt"
	"strconv"
)

func main() {

	/* ******************* 	Convert to String ******************* */

	fmt.Println(`
******************* Convert to String *******************`)
	fmt.Println(`
	******************* Float to String *******************
	`)

	//Dato tipo float
	floatVar := 33.11
	//Imprimimos por consola
	fmt.Printf("type: %T, value %f \n", floatVar, floatVar)

	//Convertimos con Sprintf el valor tipo float a string

	//%f --> variable de tipo flotante
	floatStrVar := fmt.Sprintf("%f", floatVar) // 33.110000
	//%.4f --> formateamos el dato
	floatStrVar2 := fmt.Sprintf("%.4f", floatVar) // 33.1100

	fmt.Printf("type: %T, value %s \n", floatStrVar, floatStrVar)
	fmt.Printf("type: %T, value %s \n", floatStrVar2, floatStrVar2)

	fmt.Println(`
	******************* Int to String *******************
	`)

	//Dato tipo int
	intVar := 22
	//%d --> variable tipo dato entero
	fmt.Printf("type: %T, value %d \n", intVar, intVar)

	intStrVar := fmt.Sprintf("%d", intVar)
	fmt.Printf("type: %T, value %s \n", intStrVar, intStrVar)

	//Agrego una variable de tipo int (23)
	intStrVar23 := fmt.Sprintf("%d - %d", intVar, 23)
	fmt.Printf("type: %T, value %s \n", intStrVar23, intStrVar23)

	fmt.Println(`
	******************* Int to String - package strconv *******************
	`)

	// Convertir en String utilizando package "strconv"

	intVar2 := 342
	fmt.Printf("type: %T, value %d \n", intVar2, intVar2)

	//La función Itoa recibe el valor numérico
	intStrVar2 := strconv.Itoa(intVar2)
	fmt.Printf("type: %T, value %s \n", intStrVar2, intStrVar2)

	/* ******************* 	Convert to Number ******************* */

	fmt.Println(`
******************* Convert to Number *******************`)
	fmt.Println(`
	******************* String to Int *******************
	`)

	//Dato tipo string
	//err nos arrojará error en el caso, por ejemplo, que se envíe una letra
	//La función Atoi recibe un string y lo convierte a int
	strIntVar, err := strconv.Atoi("15")
	//Imprimimos por consola
	//%d por que es un numérico
	fmt.Printf("type: %T, value %d, err: %v \n", strIntVar, strIntVar, err)
	//err = nil (nulo / null)

	//Error
	//Value 0 por que es el valor por defecto de los numéricos
	strIntVarT, err := strconv.Atoi("15T")
	fmt.Printf("type: %T, value %d, err: %v \n", strIntVarT, strIntVarT, err)

	//ParseInt --> Recibe un string
	strIntVar3, err := strconv.ParseInt("10", 10, 64) // ("valor str", "base (10 - decimal)", "tamaño (bits)")
	fmt.Printf("type: %T, value %d \n", strIntVar3, strIntVar3)

	//Error
	//strIntVarT, err := strconv.ParseInt("10T", 10, 64) // ("valor str", "base (10 - decimal)", "tamaño (bits)")
	//fmt.Printf("type: %T, value %d, err: %v \n", strIntVarT, strIntVarT, err)

	//En caso de tener un valor que no querramos utilizar ponemos "_" para omitirlo
	//para que GO no tire error

	strIntBin, _ := strconv.ParseInt("111", 2, 64) // ("valor str", "base (2 - binario)", "tamaño (bits)")
	fmt.Printf("type: %T, value %d \n", strIntBin, strIntBin)

	//Octal --> {0,1,2,3,4,5,6,7}
	strIntOctal, _ := strconv.ParseInt("11", 8, 64)
	fmt.Printf("type: %T, value %d \n", strIntOctal, strIntOctal)

	//Hexadecimal --> {0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F}
	strIntHexa, _ := strconv.ParseInt("A", 16, 64)
	fmt.Printf("type: %T, value %d \n", strIntHexa, strIntHexa)

	fmt.Println(`
	******************* String to Float *******************
	`)

	strFloatVar, err := strconv.ParseFloat("-11.2", 64) //("valor str", "tamaño")
	fmt.Printf("type: %T, value %f, err: %v \n", strFloatVar, strFloatVar, err)

}
