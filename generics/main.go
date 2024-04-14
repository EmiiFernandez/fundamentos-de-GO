package main

import (
	"fmt"

	"cmp" // Importa el paquete "cmp" que contiene la interfaz "Ordered" para la comparación ordenada
)

// La función main es la función principal donde se ejecuta el programa
func main() {
	v1 := []float64{1.3, 5.45, 45, 12.223, 6.92, 78.102}
	v2 := []int32{9, 23, 1, 23, 8, 98}

	fmt.Println(`
	---------------- Generics: Suma del mismo tipo ----------------
	`)
	fmt.Println(Sum01(v1)) // Suma de un slice de float64
	fmt.Println(Sum01(v2)) // Suma de un slice de int32

	fmt.Println(`
	---------------- Generics: Suma del mismo tipo con interface ----------------
	`)
	fmt.Println(Sum02(v1)) // Suma de un slice de float64 usando una interfaz
	fmt.Println(Sum02(v2)) // Suma de un slice de int32 usando una interfaz

	fmt.Println(`
	---------------- Generics: Diferentes tipos pero ambos valores deben ser del mismo tipo ----------------
	`)
	anyType(1, 8)     // Función que acepta cualquier tipo
	anyType("1", "A") // Función que acepta cualquier tipo
	//anyType(2, "as") // Ambos valores deben ser del mismo tipo, esta línea dará un error de compilación

	fmt.Println(`
	---------------- Generics: Diferentes tipos ----------------
	`)
	anyTypeTwo(5, "C") // Función que acepta dos tipos diferentes

	fmt.Println(`
	---------------- Generics: comparable (Igual o distinto) ----------------
	`)
	comparableType(2, 2) // Función que compara dos valores del mismo tipo
	comparableType(3, 5) // Función que compara dos valores del mismo tipo, en este caso, son diferentes

	fmt.Println(`
	---------------- Generics: cmp (interfaz compara) ----------------
	`)
	ordenedValues(2, 4) // Función que realiza operaciones con tipos comparables

	fmt.Println(`
	---------------- Generics: Slice personalizado ----------------
	`)
	csInt := CustomSlice[int]{1, 2, 3, 4} // Creación de un slice personalizado de tipo int
	fmt.Println(csInt)
	csStg := CustomSlice[string]{"a", "b", "c"} // Creación de un slice personalizado de tipo string
	fmt.Println(csStg)

	fmt.Println(`
	---------------- Generics: Mínimo entre dos números ----------------
	`)
	x, y := Point(5), Point(2)   // Creación de dos puntos
	fmt.Println(MinNumber(x, y)) // Encuentra el mínimo entre los dos puntos

	fmt.Println(`
	---------------- Generics: Estructura genérica ----------------
	`)
	fd := MyGenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}} // Creación de una estructura genérica con MyFirstData
	fd.Data.PrintOne()                                                        // Llamada al método PrintOne de MyFirstData

	sd := MyGenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}} // Creación de una estructura genérica con MySecondData
	sd.Data.PrintTwo()                                                          // Llamada al método PrintTwo de MySecondData
}

// Sum01 es una función genérica que suma elementos de un slice del mismo tipo
func Sum01[N int | int32 | int64 | float64 | float32](v []N) N {
	var total N

	for _, tV := range v {
		total += tV
	}
	return total
}

// Number es una interfaz que define varios tipos numéricos
type Number interface {
	int | int32 | int64 | float64 | float32
}

// Sum02 es una función genérica que suma elementos de un slice del mismo tipo utilizando una interfaz
func Sum02[N Number](v []N) N {
	var total N

	for _, tV := range v {
		total += tV
	}
	return total
}

// anyType es una función que acepta cualquier tipo de valores
func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

// anyTypeTwo es una función que acepta dos tipos de valores diferentes
func anyTypeTwo[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

// comparableType es una función genérica que compara dos valores del mismo tipo
// comparable solo permite consultar si el valor es igual o distinto
func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2) // Verifica si los dos valores son diferentes
}

// utilizo package go get goland.org/ex/exp
func ordenedValues[N cmp.Ordered](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 > v2)
}

// CustomSlice es una estructura genérica que representa un slice con un tipo específico de elementos
type CustomSlice[V int | string] []V

// MinNumber es una función genérica que devuelve el menor de dos números
func MinNumber[T N1](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Point es un tipo de dato entero
type Point int

// ~ se utiliza para indicar un tipo concreto dentro de una interfaz genérica. De esta forma podemos usar el Point
type N1 interface {
	~int | ~int32 | ~int64 | ~float64 | ~float32
}

// MyGenericStruct es una estructura genérica que contiene un valor y una estructura de datos
type MyGenericStruct[D any] struct {
	StrValue string
	Data     D
}

// MyFirstData es una estructura de datos vacía
type MyFirstData struct{}

// MySecondData es otra estructura de datos vacía
type MySecondData struct{}

// PrintOne es un método de MyFirstData que imprime "Print first"
func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

// PrintTwo es un método de MySecondData que imprime "Print second"
func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}
