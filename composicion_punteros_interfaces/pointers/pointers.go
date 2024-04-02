package pointers

import (
	"fmt"
)

/*
 Un puntero es una variable que almacena la dirección de memoria de otra variable. En lugar de contener el valor real de una variable, un puntero contiene la ubicación donde se almacena el valor de esa variable en la memoria.

 Ejemplo:

	Address		Value
	0x00001		2313
	0x00002		"test"
	0x00003		true
	0x00004		342.123

Para declarar un puntero se realiza:

var myPnt *int --> agregamos el puntero(*)

Ejemplo:

myVal := 2312
var myValPnt *int
myValPnt = &myVal (asigno el ponter a la address &myVal)

		&var		var
myVal	0x00001		2312
myVal	0x00002		0x00001		2312

*myValPnt --> accedemos al valor donde esta apuntando
*/

// Definir la función increment que recibe un parámetro por valor
func Increment(val int) {
	// Imprimir la dirección de memoria de la variable val
	fmt.Println("Valor de var antes de función increment: ", val)
	fmt.Println("Dirección de memoria de var de función increment:", &val)

	// Incrementar el valor de val en 1
	val++

	// Imprimir el valor de val para verificar si ha sido modificado por la función increment
	fmt.Println("Valor de var después de llamar a increment:", val)
	fmt.Println("Dirección de memoria de var después de llamar a increment:", &val)

	// No se está devolviendo nada aquí, ya que la modificación se realiza en una copia local de la variable val
	// Por lo tanto, no afecta a la variable original val
}

/*
Imprimir la dirección de memoria del entero apuntado por el puntero val y luego incrementa el valor de ese entero en 1

*val++: Incrementa el valor del entero al que apunta el puntero val en 1. Aquí, el operador * se usa para desreferenciar el puntero y acceder al valor almacenado en la dirección de memoria apuntada por val
*/
func IncrementP(val *int) {
	fmt.Println("Dirección de memoria de val dentro de IncrementP:", val) // Imprime la dirección de memoria de val
	fmt.Println("Valor de val dentro de IncrementP:", *val)               // Imprime el valor de val (el valor al que apunta el puntero)
	*val++                                                                // Incrementa el valor de la variable a la que apunta el puntero val
}
