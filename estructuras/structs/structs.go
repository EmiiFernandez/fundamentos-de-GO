package structs

import (
	"encoding/json"
	"fmt"
)

/*
Las estructuras son un conjunto de campos ("variables")
y pueden contener metódos

Ej:
Struct: User
-------------------------
-First Name | Last Name -   ---- GetName (metodo)
-Age		| Email     -   ---- GetAge
-------------------------

En GO las estructuras se implementan:

(type-palabra_reservada- nombre_estructura struct-palabra_reservada)
type name_of_struct struct {
	field1 data_type1 (nombre_de_campo tipo_de_dato )
	field2 data_type2
}

Ej:

Si la primera letra del nombre de la estructura es mayúscula será pública
si es minúscula será privado

*/

type Person struct {
	Name    string
	Age     int
	Address string
	DNI     string
	Gender  string
}

/*Dentro de las estructuras podemos utilizar Tags (etiquetas)  los cuales se utilizan para agregar una definición en un campo. Uno de los tags más utilizado es el json

Los campos que comiencen en mayúscula serán públicos.
*/

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address,omitempty"`
	Age      uint8  `json:"age,omitempty"`
	LastName string `json:"last_name"`
}

//omitempty --> al pasarlo al JSON hace que los datos vacios no se muestren

//Agrego metódos

/*
func (palabra reservada) ("variable que representa la estructura" "estructura") "nombre del metodo" {}
*/
func (u User) Display() {
	v, err := json.Marshal(u)
	fmt.Println(err)
	fmt.Println(string(v))
}

// Metodo modificar nombre

func (u *User) SetName(name string) {
	u.Name = name
}

// Debemos agregar a la estructura un * (*User) para modificar la estructura original
//Si no agrego el * solo se realizaría el cambio al ejecutar el metódo pero no la estructura original
