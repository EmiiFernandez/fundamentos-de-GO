package main

import (
	"encoding/json"
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

	fmt.Println(`
	******************* Tags & JSON *******************
	`)

	user := structs.User{
		ID:       123,
		Name:     "Nahuel",
		LastName: "García",
	}

	fmt.Println(user)

	fmt.Println(`
	******************* encoding/json *******************
	`)

	/*
		Marshal retorna una cadena de bytes
	*/
	v, err := json.Marshal(user)
	fmt.Println(err)
	fmt.Println(v)         // v--> retorna la cadena de bytes
	fmt.Println(string(v)) //Conversión de bytes a string. Formato JSON
}
