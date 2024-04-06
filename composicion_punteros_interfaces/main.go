package main

import (
	"composicion_punteros_interfaces/interfaces"
	"composicion_punteros_interfaces/interfaces/vehicles"
	"composicion_punteros_interfaces/pointers"
	"fmt"
)

type MyStruct struct {
	ID   int
	Name string
}

// Estructura SIN puntero
func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p \n", &m)
}

// Estructura con puntero
func (m MyStruct) SetP(name string) {
	fmt.Printf("address: %p \n", &m)
}

func main() {
	fmt.Println("Main")
	/* ******************* 	Composición, punteros e interfaces ******************* */
	fmt.Println(`
	******************* Ejemplo sin y con punteros *******************
	`)

	// Declarar una variable
	myVar := 30

	fmt.Println("Dirección de memoria de var fuera de la función increment: ", &myVar)
	// Llamar a la función increment y pasarle la variable por valor

	fmt.Println(`
	******************* Punteros *******************
	`)

	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println("Dirección memoría variable i: ", &i)
	fmt.Println("Puntero: ", iP)
	fmt.Println("Dirección memoría del puntero: ", &iP)
	fmt.Println("Variable i: ", i)
	fmt.Println("Valor del puntero: ", *iP)

	//Modifico el valor de iP
	*iP = 1

	//Cuando modifico uno se modifica el otro
	fmt.Printf("value i: %d, value pointer i: %d\n", i, *iP)

	fmt.Println()

	myVar = 30
	fmt.Printf("address: %p, value: %d\n", &myVar, myVar)

	fmt.Println()

	pointers.Increment(myVar)

	fmt.Println("Dirección de memoria de var fuera de la función increment: ", &myVar)

	fmt.Println()

	pointers.IncrementP(&myVar) // Llama a la función IncrementP pasando un puntero a myVar

	fmt.Println("Nuevo valor de myVar fuera de IncrementP:", myVar)

	fmt.Println(`
	------- Update Slice --------
	`)

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("address: %p, values: %v \n", mySlice, mySlice)
	fmt.Printf("addr1: %p, addr2: %p, addr3: %p \n", &mySlice[0], &mySlice[1], &mySlice[2])

	pointers.UpdateSlice(mySlice)

	fmt.Println(`
	------- Puntero de la estructura --------
	`)

	// Esta variable será un puntero de myStruct
	myStruct := &MyStruct{ID: 1234, Name: "Test"}
	// Imprime la dirección de memoria del puntero myStruct
	fmt.Printf("address: %p \n", myStruct)
	// Imprime los campos ID y Name de la estructura
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)

	// Llama a la función updateStruct pasando el puntero myStruct
	updateStruct(myStruct)
	// Imprime los campos ID y Name de la estructura después de llamar a updateStruct
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)
	fmt.Println()

	//Prueba de metodos
	fmt.Println("Estructura sin Pointer")
	myStruct.Set("test method")
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)
	fmt.Println()
	fmt.Println("Estructura con Pointer")
	myStruct.SetP("test method pointer")
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)

	fmt.Println(`
	******************* Interfaces *******************
	`)

	interfaces.Display(123)
	interfaces.Display("hola")
	interfaces.Display(true)
	interfaces.Display("123.78")

	fmt.Println(`
	******************* Interface Ejemplp *******************
	`)

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "PLANE", "MOTORCYCLE", "TRUCK", "CAR", "PLANE"}

	// variable para almacenar la distancia total
	var d float64

	// Iteramos sobre el slice de tipos de vehículos
	for _, v := range vArray {
		// Imprimimos el tipo de vehículo
		fmt.Printf("Vehicle: %s \n", v)

		// Creamos una nueva instancia del vehículo usando la función New del paquete vehicles
		veh, err := vehicles.New(v, 400)
		if err != nil {
			// Si ocurre un error al crear el vehículo, lo mostramos y pasamos a la siguiente iteración
			fmt.Println("Error: ", err)
			fmt.Println()
			continue // Continúa con la siguiente iteración del ciclo
		}

		// Calculamos la distancia recorrida por el vehículo
		distance := veh.Distance()

		// Imprimimos la distancia recorrida por el vehículo
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()

		// Sumamos la distancia recorrida por el vehículo a la distancia total
		d += distance
	}

	// Imprimimos la distancia total recorrida por todos los vehículos
	fmt.Printf("Total distance: %.2f\n", d)

}

// Recibe un puntero de la estructura
func updateStruct(v *MyStruct) {
	// Imprime la dirección de memoria del puntero recibido
	fmt.Printf("address in function: %p\n", v)
	// Actualiza los campos de la estructura
	v.ID = 999
	v.Name = "Update"
}
