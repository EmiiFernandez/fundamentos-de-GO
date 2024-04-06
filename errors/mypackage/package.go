package mypackage

import "fmt"

// La función Run es responsable de manejar cualquier pánico que pueda ocurrir dentro de ella
func Run() {
	defer func() {
		println("end my function") // Se ejecuta al final de la función
		r := recover()             // Recupera cualquier pánico que ocurra dentro de la función
		if r != nil {
			fmt.Println("error in run: ", r) // Maneja y muestra el error recuperado
		}
	}()
	panic("panic in run function") // Genera un pánico dentro de la función
}
