package vehicles

import (
	"fmt"
)

// Ejemplo de polimorfismo utilizando interfaces en Go

// Interface Vehicle define el comportamiento común para todos los tipos de vehículos.
type Vehicle interface {
	// Método Distance devuelve la distancia recorrida por el vehículo.
	Distance() float64
}

// Constantes para los tipos de vehículos
const (
	CarVehicles        = "CAR"
	MotorcycleVehicles = "MOTORCYCLE"
	TruckVehicles      = "TRUCK"
)

// Función New crea una instancia del vehículo especificado.
func New(vehicle string, time int) (Vehicle, error) {
	switch vehicle {
	case CarVehicles:
		return &Car{Time: time}, nil
	case MotorcycleVehicles:
		return &Motorcycle{Time: time}, nil
	case TruckVehicles:
		return &Truck{Time: time}, nil
	}
	return nil, fmt.Errorf("Vehicle '%s' doesn't exist", vehicle)
}

// Definición de las estructuras de los diferentes tipos de vehículos.

// Car representa un vehículo de tipo automóvil.
type Car struct {
	Time int // Tiempo en minutos
}

// Motorcycle representa un vehículo de tipo motocicleta.
type Motorcycle struct {
	Time int // Tiempo en minutos
}

// Truck representa un vehículo de tipo camión.
type Truck struct {
	Time int // Tiempo en minutos
}

// Métodos para calcular la distancia recorrida por cada tipo de vehículo.

// Distance calcula la distancia recorrida por un automóvil.
func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60) // Se asume una velocidad promedio de 100 km/h
}

// Distance calcula la distancia recorrida por una motocicleta.
func (m *Motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60) // Se asume una velocidad promedio de 120 km/h
}

// Distance calcula la distancia recorrida por un camión.
func (t *Truck) Distance() float64 {
	return 70 * (float64(t.Time) / 60) // Se asume una velocidad promedio de 70 km/h
}
