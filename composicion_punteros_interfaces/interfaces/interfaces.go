package interfaces

import (
	"fmt"
)

/*
Interfaces en GO es un conjunto de métodos abstractos. Una interfaz define un comportamiento que un tipo puede tener.

La interfaz define qué debe hacer un tipo, pero no cómo debe hacerlo. Esto permite que diferentes tipos implementen el mismo conjunto de métodos de manera diferente, pero aún así puedan ser tratados de la misma manera a través de la interfaz (Polimorfismo).

//Ejemplo
fmt.Printf("my values are %d, %s and %f \n", val1, val2, val3)

//Los argumentos se pasan como interface{} es para que puedan ser de cualquier tipo.

func Printf(format string, a ...interface{}) (n int, err error {
	return Fprintf(os.Stdout, format, a...)
})
*/

// Al agregarle interface{} podemos enviarle cualquier valor (string, int, etc)
func Display(value interface{}) {
	fmt.Println(value)
}
