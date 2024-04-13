package main

import (
	"fmt"
	"time"
)

/*
 Goroutines son funciones que se ejecutan concurrentemente con otras goroutines. Son ligeros hilos de ejecución que son administrados por el runtime de Go, permitiendo la ejecución concurrente y paralela en programas.
*/

func main() {
	//Se ejecuta A y luego B de manera secuencial (sin goroutines)
	fmt.Println(`
	---------------- SIN Goroutines ----------------
	`)

	myProcess("A")
	myProcess("B")

	//Se ejecuta ambos procesos en simultáneo mediante goroutines
	fmt.Println(`
	---------------- CON Goroutines ----------------
	`)

	// Lanzamos dos goroutines concurrentes, cada una ejecutando la función myProcess
	go myProcess("A")
	go myProcess("B")

	// Esperamos 3 segundos para que las goroutines tengan tiempo de ejecutarse
	// En un programa real, este tiempo de espera podría variar dependiendo de las necesidades del programa
	time.Sleep(3 * time.Second)
}

// La función myProcess simula un proceso que realiza alguna tarea
// En este caso, simplemente imprime un mensaje indicando el proceso y un número incremental
func myProcess(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(500 * time.Millisecond) // Simulamos un retraso de 500 milisegundos (0.5 segundos)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
