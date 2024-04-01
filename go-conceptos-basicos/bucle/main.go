package main

import (
	"fmt"
)

func main() {

	/* ******************* 	Bucle ******************* */

	fmt.Println(`
	******************* Bucle for *******************
	`)

	//En GO solo tenemos el bucle for. No existe el while

	//Sumatoria 1 sin iterar
	sum := 0
	//sum++
	fmt.Println("Sum = ", sum)

	for i := 0; i < 10; i++ {
		//fmt.Println(i) //Imprimo los valores de i
		sum++
	}
	fmt.Println("Resultado de sum luego del bucle for: ", sum)

	//Ejecutamos x veces mientras que sum sea menor a 1000
	for sum < 1000 {
		sum++
	}
	fmt.Println(sum) // Sale del bucle

	sum = 0
	//for sin condición
	for {
		if sum > 1000 {
			break //rompo el bucle con el if
		}
		sum++
	}
	fmt.Println(sum)
}
