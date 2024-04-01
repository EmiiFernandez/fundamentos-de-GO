package main

import (
	"fmt"
)

func main() {

	/* ******************* 	Slices ******************* */

	fmt.Println(`
	******************* Slices *******************
	`)

	//A diferencia de los arrays pueden modificar su tamaño en tiempo de ejecución
	//Podemos obtener un slice en base a otro slice o un array, por ej. extraer una porción de otro slice. Tener cuidado que el cambio se verá reflejado tanto en el nuevo slice como en el slice o array original
	//Toma la misma dirección de memoría del valor original (address)

	myArrayVar := [5]int{2, 6, 9, 10, 16}
	fmt.Println("array: ", myArrayVar, " - lenght: ", len(myArrayVar))

	mySliceVar := []int{} // inicializo siempre con [] vacios

	fmt.Println("slice: ", mySliceVar, " - lenght: ", len(mySliceVar))

	//agregar datos al slice
	//No permite agregar un array a un slice con append
	mySliceVar = append(mySliceVar, 12, 34, 54)
	fmt.Println("slice: ", mySliceVar, " - lenght: ", len(mySliceVar))

	mySliceVar = append(mySliceVar, 31, 12)
	fmt.Println("slice: ", mySliceVar, " - lenght: ", len(mySliceVar))

	fmt.Println("array del index 2:4: ", myArrayVar[2:4])
	mySliceVar2 := myArrayVar[2:4]
	fmt.Println("slice de myArrayVar[2:4]: ", mySliceVar2, " - lenght: ", len(mySliceVar2))

	fmt.Println(`
	******************* Modifico slices *******************
	`)

	fmt.Println("array del index 2: ", myArrayVar[2])
	fmt.Println("slice del index 0: ", mySliceVar2[0])
	mySliceVar2[0] = 19
	fmt.Println("array del index 2 modificado: ", myArrayVar[2], " - address: ", &myArrayVar[2])
	fmt.Println("array del index 0: ", mySliceVar2[0], " - address: ", &mySliceVar2[0])

	mySliceVar3 := mySliceVar[:4] //desde index 0
	fmt.Println(mySliceVar3)

	mySliceVar4 := mySliceVar[2:] //desde el index 2 hasta el final
	fmt.Println(mySliceVar4)

	fmt.Println(`
	******************* Inicializar slices con make *******************
	`)

	mySliceVar5 := make([]int, 80) //Pasamos el tipo de dato y la cantidad de elementos.
	fmt.Println("slice: ", mySliceVar5, " - lenght: ", len(mySliceVar5))

	fmt.Println(`
	******************* Inicializar slices con valores *******************
	`)

	mySliceVar6 := []int{1, 2, 6, 11, 20, 5, 1, 0} //Pasamos el tipo de dato y la cantidad de elementos.
	fmt.Println("slice: ", mySliceVar6, " - lenght: ", len(mySliceVar6))

}
