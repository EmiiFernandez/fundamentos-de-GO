package main

import (
	"fmt"
)

func main() {

	/* ******************* 	Operators ******************* */

	fmt.Println(`
******************* Operators *******************`)

	/*
		   Los operadores nos permiten realizar más de una validación

		   Los operadores incluyen:
		   	*la negación lógica unaria (!),
		   	*AND lógico binario (&), OR (|) y OR exclusivo (^)
		   	*y los AND (&&) y OR (||) lógicos condicionales binarios.


		  Operandos		AND	OR
		   	V	V	|	V	V
		   	V	F	|	F	V
		   	F	V	|	F	V
		   	F	F	|	F	F

			Operandos	NOT
				V	|	F
				F	|	V

	*/

	fmt.Println(`
	******************* Validaciones *******************
	`)

	yearsOld := 33
	fmt.Println("El valor a validar es: ", yearsOld)
	fmt.Println("¿yearsOld es mayor a 30? ", yearsOld > 30)          //true
	fmt.Println("¿yearsOld es menor a 33? ", yearsOld < 33)          //false
	fmt.Println("¿yearsOld es menor o igual a 33? ", yearsOld <= 33) //true
	fmt.Println("¿yearsOld es mayor o igual a 40? ", yearsOld >= 40) //false
	fmt.Println("¿yearsOld es igual a 33? ", yearsOld == 33)         // true

	fmt.Println(`
	******************* Validaciones operador OR (||) *******************
	`)

	fmt.Println("yearsOld es menor a 33 o igual a 33?", yearsOld < 33 || yearsOld == 33) // false - true = true
	fmt.Println("yearsOld es menor a 33 o igual a 34?", yearsOld < 33 || yearsOld == 34) // false - false = false
	fmt.Println("yearsOld es menor a 40 o igual a 33?", yearsOld < 40 || yearsOld == 33) // true - true = true

	fmt.Println(`
	******************* Validaciones operador AND (&&) *******************
	`)

	fmt.Println("yearsOld es menor a 33 o igual a 33?", yearsOld < 33 && yearsOld == 33) // false - true = false
	fmt.Println("yearsOld es menor a 33 o igual a 34?", yearsOld < 33 && yearsOld == 34) // false - false = false
	fmt.Println("yearsOld es menor a 40 o igual a 33?", yearsOld < 40 && yearsOld == 33) // true - true = true

	fmt.Println(`
	******************* Validaciones operador NOT (!) *******************
	`)

	fmt.Println("true: ", true)                               // true
	fmt.Println("NOT de true: ", !true)                       //false
	fmt.Println("yearsOld menor a 40: ", yearsOld < 40)       // true
	fmt.Println("!(yearsOld menor a 40): ", !(yearsOld < 40)) // true

	fmt.Println(`
	******************* Validaciones operadores combinados *******************
	`)

	//Primero se ejecuta primero el AND y luego el OR
	//						false				true
	fmt.Println("yearsOld < 25 && yearsOld == 33 || yearsOld < 40: ", yearsOld < 25 && yearsOld == 33 || yearsOld < 40) // false AND true OR true = true
	//Para ejecutar primero el OR hay que agregarlo entre parentesis
	fmt.Println("yearsOld < 25 && (yearsOld == 33 || yearsOld < 40): ", yearsOld < 25 && (yearsOld == 33 || yearsOld < 40)) // false AND (true OR true) = false

}
