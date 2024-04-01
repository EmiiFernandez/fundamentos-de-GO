package main

import (
	"fmt"
	"funciones/function"
)

func main() {

	/* ******************* 	Funciones ******************* */

	fmt.Println(`
	******************* Funciones *******************
	`)

	/*  Para crear una funcion utilizo func (palabra reservada), nombre de la función (nombre_valores tipo_de_dato) tipo_de_dato (el tipo de dato que retorna la función)
	Ejemplo:
	func myFunction (myValue int64, myValue2 int64) int64 {
		scope de la función
	}*/
	var myIntVar int64
	var myIntVar1 int64 = 30
	var myIntVar2 int64 = 50
	var myIntVar3 int64 = 90

	fmt.Println(`
	******************* Desde el main.go *******************
	`)

	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v \n", myIntVar, myIntVar, &myIntVar)
	fmt.Println()

	fmt.Println(`
	******************* Importando función *******************
	`)

	//utilizando función importada
	function.Display(myIntVar)
	function.Display(myIntVar1)
	function.Display(myIntVar2)
	function.Display(myIntVar3)

	fmt.Println(`
	******************* Función que retorna int *******************
	`)

	v := function.Add(4, 2)
	fmt.Println("La suma 4 + 2 es: ", v)

	fmt.Println(`
	******************* Función repetir string *******************
	`)

	function.RepeatString(10, "LA")
	function.RepeatString(10, "LE")

	fmt.Println(`
	******************* Función con switch *******************
	`)

	value, err := function.Calc(function.SUM, 20.12, 34)
	fmt.Println("Suma - value: ", value, " - error: ", err)
	value2, err := function.Calc(function.SUB, 20.12, 34)
	fmt.Println("Resta - value: ", value2, " - error: ", err)
	value3, err := function.Calc(function.DIV, 20.12, 34)
	fmt.Println("División - value: ", value3, " - error: ", err)
	value4, err := function.Calc(function.DIV, 20.12, 0)
	fmt.Println("División - value: ", value4, " - error: ", err)
	value5, err := function.Calc(function.MUL, 20.12, 34)
	fmt.Println("Multiplicación - value: ", value5, " - error: ", err)
	value6, err := function.Calc(function.TEST, 20.12, 34)
	fmt.Println("Test error - value: ", value6, " - error: ", err)

}
