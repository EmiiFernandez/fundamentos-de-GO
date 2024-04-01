package main

import (
	"fmt"
)

func main() {
	//Agregamos variables dentro de un scope para que solo existan
	//dentro de ese scope.
	//Debemos agregarlas entre llaves {}

	myOtherScopeVariable := 50

	{ //Primer scope
		myScopeVariable := 40
		//Los niveles de scope siguientes pueden utilizar las variables
		//que se encuentran por encima de la llave
		{
			fmt.Println("Mi variable fuera del scope es: ", myOtherScopeVariable)
			fmt.Println("Mi variable dentro del primer scope es: ", myScopeVariable)
		}
		//fmt.Println("Mi variable fuera del scope es: ", myOtherScopeVariable)
		///fmt.Println("Mi variable dentro del primer scope es: ", myScopeVariable)
	}

	//fmt.Println("Mi variable fuera del scope es: ", myScopeVariable) //.\main.go:17:50: undefined: myScopeVariable
	fmt.Println("Mi variable fuera del scope es: ", myOtherScopeVariable)

	{ //Segundo scope
		fmt.Println("Mi variable en segundo scope es: ", myOtherScopeVariable)
		//fmt.Println("Mi variable del primer scope: ", myScopeVariable)
		//No se puede usar myScopeVariable por que se encuentra en otro scopes
	}
}
