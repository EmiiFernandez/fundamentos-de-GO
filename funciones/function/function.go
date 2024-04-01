package function

import (
	"errors"
	"fmt"
)

type Operation int // Defino un tipo de dato que será representado por un int

// Generamos un Enum utilizando el tipo de dato Operation que será un int
const (
	SUM Operation = iota
	SUB
	DIV
	MUL
	TEST
)

//TODA FUNCIÓN/VARIABLES/METODOS QUE COMIENCE CON MAYÚSCULA ES PÚBLICO
//SE PUEDE ACCEDER DESDE AFUERA

//SI ARRANCA EN MINÚSCULA SERÁ PRIVADA. SOLO PODRÁN ACCEDER LOS QUE SE
//ENCUENTREN DENTRO DEL MISMO PACKAGE

func Display(myValue int64) {
	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v \n", myValue, myValue, &myValue)
	fmt.Println()

}

// recibimos dos variables y retonarnará tipo de dato int
func Add(x int, y int) int {
	//t := x + y
	return x + y
}

// Repetirá x veces (increment) el valor (value)
// Por pantalla se verá solo el valor repetido
func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value) // sin salto de linea
	}

	fmt.Println()
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y no puede ser cero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}
	return 0, errors.New("Ha ocurrido un error")
}

/* PRIVADO
func display(myValue int64) {
	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v \n", myValue, myValue, &myValue)
	fmt.Println()
}*/
