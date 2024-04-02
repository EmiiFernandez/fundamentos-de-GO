package function

//Como estamos dentro del package de function accedemos a Operation
//sin necesidad de importar
//Ingresamos que operación queremos realizar
//Vamos a retornar una func que recibirá dos floats y devolverá un float64
func FactoryOperation(op Operation) func(float64, float64) float64 {
	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case MUL:
		return mul
	case DIV:
		return div
	}
	return nil //en caso que no se haga ninguna de las operaciones
}

//FUNCIONES PRIVADAS.

// Pasamos una operación y retorna una función
func sub(a, b float64) float64 {
	return a - b
}

func sum(a, b float64) float64 {
	return a + b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
