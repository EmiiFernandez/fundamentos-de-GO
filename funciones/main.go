package main

import (
	"codigo/funciones/function"
	"fmt"
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

	fmt.Println(`
	******************* Función que recibe un valor retorna dos variables *******************
	`)

	xVal, yVal := function.Split(40)
	fmt.Println("Valor x: ", xVal, "Valor y: ", yVal)

	fmt.Println(`
	******************* Función que suma x valores *******************
	`)

	val2 := function.MSum(1, 2, 3, 1, 2, 3, 4)
	fmt.Println(val2)

	val3 := function.MSum(1, 2, 3, 1, 2, 3, 4, 90, 100)
	fmt.Println(val3)

	fmt.Println(`
	******************* Función combinada *******************
	`)

	mOperVal, err := function.MOperations(function.SUM, 4, 4, 4, 4, 10, 80)
	fmt.Println("SUMA: value: ", mOperVal, " - error: ", err)
	mOperVal2, err := function.MOperations(function.SUB, 4, 4, 4, 80)
	fmt.Println("RESTA: value: ", mOperVal2, " - error: ", err)
	mOperVal3, err := function.MOperations(function.MUL, 4, 4, 10, 80, 6)
	fmt.Println("MULTIPLICACIÓN: value: ", mOperVal3, " - error: ", err)
	mOperVal4, err := function.MOperations(function.DIV, 4, 4, 4, 4, 10, 80)
	fmt.Println("DIVISIÓN: value: ", mOperVal4, " - error: ", err)
	mOperVal5, err := function.MOperations(function.DIV, 0, 0)
	fmt.Println("DIVISIÓN por 0: value: ", mOperVal5, " - error: ", err)

	fmt.Println(`
	******************* Función factory *******************
	`)

	factOpFunc := function.FactoryOperation(function.SUM)
	fmt.Println("SUMA: ", factOpFunc(10, 15))
	factOpFuncSub := function.FactoryOperation(function.SUB)
	fmt.Println("RESTA: ", factOpFuncSub(10, 15))
	factOpFuncDiv := function.FactoryOperation(function.DIV)
	fmt.Println("DIVISIÓN: ", factOpFuncDiv(10, 15))
	factOpFuncDiv0 := function.FactoryOperation(function.DIV)
	fmt.Println("DIVISIÓN x 0: ", factOpFuncDiv0(10, 0))
	factOpFuncMul := function.FactoryOperation(function.MUL)
	fmt.Println("MULTIPLICACIÓN: ", factOpFuncMul(10, 15))
	factOpFuncMulAll := function.FactoryOperation(function.MUL)(100, 15)
	fmt.Println("MULTIPLICACIÓN: ", factOpFuncMulAll)

}
