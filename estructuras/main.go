package main

import (
	"estructuras/structs"
	"fmt"
)

func main() {
	fmt.Println("Main")
	/* ******************* 	Estructuras ******************* */

	fmt.Println(`
	******************* Structs *******************
	`)

	// Crear una nueva instancia de Person

	person := structs.Person{
		Name:    "Emilia",
		Age:     30,
		Address: "Calle 123",
		DNI:     "56369852",
		Gender:  "Femenino",
	}

	// Acceder a los campos de la estructura Person
	fmt.Println("Estructura completa: ", person)
	fmt.Println("Nombre:", person.Name)
	fmt.Println("Edad:", person.Age)
	fmt.Println("Dirección:", person.Address)
	fmt.Println("DNI:", person.DNI)
	fmt.Println("Género:", person.Gender)
}
