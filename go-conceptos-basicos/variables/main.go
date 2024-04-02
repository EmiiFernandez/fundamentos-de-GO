package main

import (
	"fmt"
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

	/*
			Convenciones para nombrar elementos

		Existen algunas convenciones que podemos utilizar para nombrar variables, funciones y otros elementos en la programación.
		*Camel Case: las palabras se escriben juntas sin espacios y cada palabra después de la primera comienza con una letra mayúscula (excepto la primera palabra).
		Ejemplo: miVariableEjemplo o calcularSalarioNeto
		*Snake Case: las palabras se escriben en mayúsculas y se separan mediante guiones bajos (underscore).
		Ejemplo: mi_variable_ejemplo o calcular_salario_neto
		Pascal Case: Es similar a Camel Case, pero en lugar de comenzar con una letra mayúscula, la primera letra de cada palabra se inicia en mayúscula.
		Ejemplo: MiVariableEjemplo o CalcularSalarioNeto
		*Kebab Case: las palabras se escriben en minúsculas y se separan mediante guiones. Aunque no es tan común en la programación, se usa en ciertos contextos, especialmente en la definición de la URLs y nombres de archivos web.
		Ejemplo:
		mi-variable-ejemplo o calcular-salario-neto
	*/
}
