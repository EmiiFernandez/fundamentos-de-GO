package main

import (
	"fmt"
)

func main() {

	/* ******************* 	Maps ******************* */

	fmt.Println(`
	******************* Maps *******************
	`)

	//Estructuras que son representadas por la estructura clave (key) - valor (value)
	//Se pueden modificar o eliminar un valor por una key
	//Si queremos obtener el valor de un elemento dentro de los maps debemos acceder a la key
	//Si intento acceder a una key que no existe devuelve vacío

	//make inicializa el maps
	//map[tipo de dato de la key]tipo de dato del value
	map1 := make(map[int]string) //Inicializa el mapa

	fmt.Println(map1) // vacio

	//asigno valores
	map1[1] = "A" // key 1 - value A
	map1[2] = "B"
	map1[3] = "C"
	map1[99] = "Z"
	map1[-5] = "a"
	fmt.Println(map1)

	//acceder a un valor por su key
	fmt.Println("El valor de la key 1 es: ", map1[1])

	map2 := make(map[int][]string) //Como value tiene un array string
	map2[1] = []string{"A", "B"}
	map2[3] = []string{"Z", "A", "9"}

	fmt.Println("El map2 es: ", map2)
	fmt.Println("El valor de la key 1 es: ", map2[1])

	//Mapa inicializado desde el comienzo
	map3 := map[int]string{
		4: "A",
		1: "N",
		8: "90",
	}

	fmt.Println(map3)

	fmt.Println(`
	******************* Modificar maps *******************
	`)

	fmt.Println("key 1: ", map1[1])
	map1[1] = "LA"
	fmt.Println("key 1 modificada: ", map1[1])

	fmt.Println(`
	******************* Eliminar maps *******************
	`)

	fmt.Println("key 2: ", map1[2])
	delete(map1, 2) //palabra reservada delete (maps, key)
	fmt.Println("key 2 eliminada: ", map1[2])

	fmt.Println(`
	******************* Buscar un valor *******************
	`)

	v, ok := map1[9] // v -> value - ok -> Si el valor existe o no
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	v, ok = map1[1] //
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	fmt.Println(`
	******************* Recorrer maps con for *******************
	`)

	for k := range map1 { // k -> key - range map1 (tamaño del maps)
		fmt.Println("key: ", k, " - value: ", map1[k])
	}

	for k, v := range map1 { // k -> key - v -> value
		fmt.Println("key: ", k, " - value: ", v)
	}

	for key, value := range map2 { // k -> key - v -> value
		fmt.Println("key: ", key)
		for _, v := range value {
			fmt.Println("Value: ", v)
		}
		fmt.Println() //Salto de linea
	}
}
