package main

import (
	"encoding/json"
	"estructuras/commerce"
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

	fmt.Println(`
	******************* Metodos *******************
	`)

	//Imprimo en formato JSON los datos de User
	user.Display()

	//Modificar nombre del usuario
	user.SetName("Azul")
	user.Display()

	fmt.Println(`
	******************* Ejemplo Commerce *******************
	`)

	//Creo producto que lleva el usuario
	p1 := commerce.Product{
		Name:  "Heladera marca sarasa",
		Price: 200000,
		Type: commerce.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "freezer", "sarasa", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Naranja",
		Price: 50,
		Type: commerce.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "fruta"},
		Count: 20,
	}

	p3 := commerce.Product{
		Name:  "Cortina",
		Price: 6000,
		Type: commerce.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	//Creo el carrito de compras
	car := commerce.NewCar(11312)
	//Agrego productos al carrito
	car.AddProducts(p1, p2, p3)

	fmt.Println("PRODUCTS CAR")
	// Imprimir el nombre y el precio de los productos comprados
	fmt.Println("Productos comprados:")
	for _, product := range car.Products {
		fmt.Printf("Nombre: %s, Precio: $%.2f, Cantidad: %d\n", product.Name, product.Price, product.Count)
	}

	//Imprimir cantidad total de productos comprados
	var total uint16
	for _, product := range car.Products {
		total += product.Count
	}
	fmt.Println("Total products: ", total)

	//Imprimir precio total
	fmt.Printf("Total Car: $ %.2f \n", car.Total())
}
