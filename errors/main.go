package main

import (
	"errors" // Importamos el paquete errors para trabajar con errores en Go
	"fmt"
)

// Definimos una estructura MyCustomErr que implementa la interfaz Error de Go
type MyCustomErr struct {
	ID string // Un campo para identificar el error
}

// Implementamos el método Error de la interfaz Error para MyCustomErr
func (m MyCustomErr) Error() string {
	return fmt.Sprintf("error with id: %s", m.ID) // Devolvemos una cadena formateada que describe el error
}

func main() {
	var err error

	// Creamos un nuevo error utilizando la función New del paquete errors
	err = errors.New("my new error")
	fmt.Println(err)
	fmt.Println(err.Error())

	// Creamos un error formateado utilizando la función Errorf del paquete fmt
	err2 := fmt.Errorf("my format err, string: %s, number: %d", "my string", 231)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	// As: Llamamos a la función TestError para obtener un error personalizado
	err3 := TestError(1)
	fmt.Println(err3)
	fmt.Println(errors.As(err3, &MyCustomErr{})) // Verificamos si el error es del tipo MyCustomErr

	// As: Llamamos a la función TestError con un valor diferente para obtener un error genérico
	err4 := TestError(4)
	fmt.Println(err4)
	fmt.Println(errors.As(err4, &MyCustomErr{})) // Intentamos verificar si el error es del tipo MyCustomErr, pero falla

	// Join: Combinamos varios errores en uno utilizando la función Join del paquete errors
	err5 := errors.Join(err, err2, err3)
	fmt.Println(err5)

	// Is: Verificamos si un error específico está contenido dentro de err5
	fmt.Println(errors.Is(err5, err2))
	fmt.Println(errors.Is(err5, err4))

	// Creamos errores envueltos para mostrar la cadena de errores
	err6 := fmt.Errorf("error1: [%w]", err)
	err7 := fmt.Errorf("error2: [%w]", err6)
	fmt.Println(err7)

	// Unwrap: Desenvolvemos los errores envueltos para obtener el error original
	fmt.Println(errors.Unwrap(err7))
	fmt.Println(errors.Unwrap(errors.Unwrap(err7)))
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err7))))
}

// TestError es una función que devuelve un error personalizado según el valor pasado
func TestError(v int) error {
	if v == 1 {
		return MyCustomErr{"4"} // Devuelve un error personalizado si v es igual a 1
	}
	return errors.New("my error") // Devuelve un error genérico en otros casos
}
