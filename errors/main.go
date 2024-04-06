package main

import (
	"codigo/errors/mypackage"
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
	fmt.Println(`
	---------------- Función New ----------------
	`)
	// Creamos un nuevo error utilizando la función New del paquete errors
	err = errors.New("my new error")
	fmt.Println(err)
	fmt.Println(err.Error())

	fmt.Println(`
	---------------- Crear error con Errorf ----------------
	`)
	// Creamos un error formateado utilizando la función Errorf del paquete fmt
	err2 := fmt.Errorf("my format err, string: %s, number: %d", "my string", 231)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	fmt.Println(`
	---------------- Función As ----------------
	`)
	// As: Llamamos a la función TestError para obtener un error personalizado
	err3 := TestError(1)
	fmt.Println(err3)
	fmt.Println(errors.As(err3, &MyCustomErr{})) // Verificamos si el error es del tipo MyCustomErr

	// As: Llamamos a la función TestError con un valor diferente para obtener un error genérico
	err4 := TestError(4)
	fmt.Println(err4)
	fmt.Println(errors.As(err4, &MyCustomErr{})) // Intentamos verificar si el error es del tipo MyCustomErr, pero falla

	fmt.Println(`
	---------------- Función Join ----------------
	`)
	// Join: Combinamos varios errores en uno utilizando la función Join del paquete errors
	err5 := errors.Join(err, err2, err3)
	fmt.Println(err5)

	fmt.Println(`
	---------------- Función Is ----------------
	`)
	// Is: Verificamos si un error específico está contenido dentro de err5
	fmt.Println(errors.Is(err5, err2))
	fmt.Println(errors.Is(err5, err4))

	fmt.Println(`
	---------------- Cadena de errores - Errorf ----------------
	`)

	// Creamos errores envueltos para mostrar la cadena de errores
	err6 := fmt.Errorf("error1: [%w]", err)
	err7 := fmt.Errorf("error2: [%w]", err6)
	fmt.Println(err7)

	fmt.Println(`
	---------------- Función Unwrap ----------------
	`)
	// Unwrap: Desenvolvemos los errores envueltos para obtener el error original
	fmt.Println(errors.Unwrap(err7))
	fmt.Println(errors.Unwrap(errors.Unwrap(err7)))
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err7))))

	/*
	   errors.New: Crear un nuevo error con un mensaje específico. Retorna un valor de tipo error. Por ejemplo:
	   err := errors.New("este es un mensaje de error")

	   errors.Is: Verificar si un error es igual a otro error. Retorna un valor booleano que indica si el error dado es igual al error especificado. Se utiliza para manejar errores específicos en el código. Por ejemplo:
	   if errors.Is(err1, err2) {
	       // manejar el error
	   }

	   errors.Join: Combinar varios errores en uno solo. Retorna un error que contiene la concatenación de los mensajes de error de los errores dados. Se utiliza para agrupar múltiples errores en un solo error. Por ejemplo:
	   err := errors.Join(err1, err2, err3)

	   errors.Unwrap: Se utiliza para desenvolver un error que ha sido envuelto previamente. Retorna el error envuelto original que se pasó a la función fmt.Errorf con el especificador %w. Se utiliza para extraer el error original de una cadena de errores envueltos. Por ejemplo:
	   originalErr := errors.Unwrap(wrappedErr)
	*/

	fmt.Println(`
	---------------- Defer / Panic / Recover ----------------
	`)

	// Define una función defer que se ejecutará al final del programa
	// Permite recuperar errores que puedan ocurrir durante la ejecución
	defer func() {
		fmt.Println("end main") // Se ejecuta al final del programa
		r := recover()          // Recupera cualquier error que se haya producido (si lo hay)
		if r != nil {
			fmt.Println("Recover in ", r) // Maneja y muestra el error recuperado en caso de panic
		}
	}()

	// Dispara un error
	// Se coloca detrás del defer para asegurarse de que el defer se ejecute antes de que se produzca el panic
	// panic("my panic")
	mypackage.Run() // Llama a una función que puede generar un panic

	fmt.Println("end main") // Se ejecuta antes de "end main"

	/*
	   panic: Se utiliza para generar un error fatal. Cuando se encuentra un panic, la ejecución del programa se detiene inmediatamente y se ejecuta cualquier función defer pendiente antes de salir del programa. El mensaje pasado a panic puede ser cualquier valor, no solo una cadena.

	   defer: Se utiliza para posponer la ejecución de una función hasta que la función que la rodea haya terminado de ejecutarse. Las funciones defer se ejecutan en el orden inverso en el que se definen, es decir, la última en ser definida será la primera en ser ejecutada.

	   recover: Se utiliza para recuperar y manejar un pánic. Se coloca dentro de una función defer. Si no hay ningún panic, recover devuelve nil. Si se llama a recover dentro de una función defer y la función circundante ha generado un panic, recover detiene la propagación del panic y devuelve el valor pasado a panic. Esto permite que el programa continúe su ejecución después de un panic y proporciona una manera de manejarlo de manera controlada.
	*/

}

// TestError es una función que devuelve un error personalizado según el valor pasado
func TestError(v int) error {
	if v == 1 {
		return MyCustomErr{"4"} // Devuelve un error personalizado si v es igual a 1
	}
	return errors.New("my error") // Devuelve un error genérico en otros casos

}
