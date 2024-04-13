package main

import (
	"fmt"
	"time"
)

/*
 Goroutines son funciones que se ejecutan concurrentemente con otras goroutines.
 Son ligeros hilos de ejecución que son administrados por el runtime de Go,
 permitiendo la ejecución concurrente y paralela en programas.

Los channels son un mecanismo de comunicación que permite la sincronización y la comunicación entre goroutines en Go.
Se pueden utilizar para enviar y recibir valores de un extremo a otro, lo que permite la coordinación entre goroutines.
*/

func main() {
	//Se ejecuta A y luego B de manera secuencial (sin goroutines)
	fmt.Println(`
	---------------- SIN Goroutines ----------------
	`)

	// Ejecución secuencial de los procesos A y B
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

	fmt.Println(`
	---------------- Goroutines + Channels ----------------
	`)

	// Lanzamos dos goroutines concurrentes, cada una ejecutando la función myProcess
	go myProcess("A")
	go myProcess("B")

	// Creamos un canal (channel) para comunicarnos entre goroutines
	myFirstChannel := make(chan string)

	// Lanzamos una goroutine adicional que utiliza un canal para enviar datos
	go myProcessWithChannel("C", myFirstChannel)

	// Recibimos el resultado enviado por la goroutine usando el canal
	result := <-myFirstChannel
	fmt.Println(result)

	// Cerramos el canal después de usarlo
	close(myFirstChannel)

	fmt.Println(`
	---------------- Goroutines + 2 Channels ----------------
	`)

	// Creamos dos canales separados para comunicarnos con diferentes goroutines
	channelA := make(chan string)
	channelB := make(chan string)

	// Lanzamos dos goroutines distintas, cada una usando su propio canal para enviar datos
	go myProcessWithChannel("D", channelA)
	go myOtherProcessWithChannel("E", channelB)

	// Recibimos los resultados enviados por las goroutines usando los canales correspondientes
	resultA := <-channelA
	fmt.Println("A: ", resultA)
	resultB := <-channelB
	fmt.Println("B: ", resultB)

	// Cerramos los canales después de usarlos
	close(channelA)
	close(channelB)
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

// La función myProcessWithChannel es similar a myProcess, pero envía un mensaje a través de un canal cuando termina
func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 5 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	// Enviamos un mensaje al canal para indicar que el proceso ha terminado
	c <- "ok"
}

// myOtherProcessWithChannel es otra función similar a myProcessWithChannel
// Pero con una duración diferente para ilustrar el uso de múltiples canales con goroutines distintas
func myOtherProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	// Enviamos un mensaje al canal para indicar que el proceso ha terminado
	c <- "ok"
}
