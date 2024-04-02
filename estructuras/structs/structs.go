package structs

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
