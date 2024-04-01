package main

import (
	"fmt"
)

func main() {

	/* ******************* 	Condicionales ******************* */

	fmt.Println(`
******************* Condicionales *******************`)

	fmt.Println(`
	******************* Condicional if *******************`)

	yearsOld := 33

	fmt.Println("El valor a validar es: ", yearsOld)

	//Si yearsOld es mayor a 18 imprime por consola, sino queda vacío
	if yearsOld > 18 {
		fmt.Printf("%d es mayor a 18 \n", yearsOld)
	}

	boolVar := true

	if boolVar {
		fmt.Println("La variable boolVar es verdadero")
	}

	if !boolVar == false {
		fmt.Println("El NOT de boolVar es false")
	}

	if boolVar {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	//puedo asignarle un valor en el mismo if
	if value := true; value {
		fmt.Println("value es verdadero")
	}

	number := 5

	if number == 1 {
		fmt.Println("uno")
	} else if number == 2 {
		fmt.Println("dos")
	} else if number == 3 {
		fmt.Println("tres")
	} else {
		fmt.Println("ninguna es válida")
	}

	fmt.Println(`
	******************* Condicional switch *******************
	`)

	fmt.Println("El valor de number es: ", number)

	switch number {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("uno")
	default: //Si ninguna se cumple
		fmt.Println("ninguna es válida")
	}

	//Puedo asignar el valor dentro del switch

	switch numberOne := 1; numberOne {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("uno")
	default: //Si ninguna se cumple
		fmt.Println("ninguna es válida")
	}

	//No agrego variables para validar en el switch, lo hacemos en el case

	switch {
	case number > 0:
		fmt.Println("positivo")
	case number < 0:
		fmt.Println("negativo")
	case number == 0:
		fmt.Println("es cero")
	}
}
