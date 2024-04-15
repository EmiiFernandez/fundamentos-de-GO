package main

import (
	"context" // Paquete para trabajar con contextos
	"fmt"
	"time"
)

func main() {
	// Creamos un contexto de fondo (background context)
	ctx := context.Background()

	// Agregamos valores al contexto utilizando context.WithValue
	ctx = context.WithValue(ctx, "my-key", "my value")
	// Agregar una cadena de texto
	ctx = context.WithValue(ctx, "my-key-int", 5)
	// agregar un entero
	viewContext(ctx)
	// Visualizar los valores del contexto

	// Creamos un nuevo contexto con un límite de tiempo (timeout)
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Llamamos a cancel() para liberar los recursos asociados al contexto cuando terminemos su uso

	// Lanzamos una goroutine que realiza un proceso con el contexto que tiene un límite de tiempo
	go myProcess(ctx2)

	// Esperamos hasta que el contexto se cancele o expire
	select {
	case <-ctx2.Done():
		fmt.Println("we've exceeded the deadline")
		fmt.Println(ctx2.Err()) // Mostramos el error que causó la cancelación o expiración del contexto

	}
}

// viewContext muestra los valores almacenados en el contexto
func viewContext(ctx context.Context) {
	fmt.Printf("my value is: %s\n", ctx.Value("my-key"))     // Obtener el valor asociado a "my-key" (cadena)
	fmt.Printf("my value is: %d\n", ctx.Value("my-key-int")) // Obtener el valor asociado a "my-key-int" (entero)
}

// myProcess realiza un proceso que puede ser cancelado si el contexto es cancelado o expira
func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ohh, time out")
			return // Salimos de la función cuando el contexto es cancelado o expira
		default:
			fmt.Println("doing some process") // Simulamos hacer algún proceso
		}
		time.Sleep(500 * time.Millisecond) // Esperamos un corto tiempo antes de verificar el contexto nuevamente
	}
}

/*
context:

Es una estructura que se utiliza para llevar valores y señales de cancelación a través de las llamadas a funciones y goroutines. Proporciona una manera de controlar la cancelación y los plazos de ejecución en programas concurrentes.

Aquí hay algunas situaciones en las que se utiliza un contexto:

*Gestión de cancelación: se puede utilizar para cancelar operaciones en curso, como solicitudes HTTP, procesamiento de datos o acceso a recursos externos, cuando una operación superior o el usuario deciden que ya no son necesarias.

*Control de tiempo de espera: puede configurarse con un límite de tiempo para operaciones que deben completarse dentro de un plazo determinado. Si la operación no se completa dentro de este tiempo, el contexto se cancela automáticamente y se notifica a las partes interesadas.

*Propagación de valores y datos: puede llevar valores asociados que pueden ser útiles para las operaciones que se realizan dentro de él. Estos valores se pueden utilizar para pasar información relevante a lo largo de una goroutines.

*Gestión de errores y comunicación entre goroutines: Los contextos proporcionan una manera estructurada de manejar errores y comunicarse entre goroutines. Pueden utilizarse para coordinar la terminación de múltiples goroutines y para propagar errores de manera eficiente.

*/
