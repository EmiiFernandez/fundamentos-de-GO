package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	p := fmt.Println

	// Paquete Time: operaciones relacionadas con el tiempo
	fmt.Println(`
	---------------- Package  Time ----------------
	`)

	// Fecha y hora actual
	now := time.Now()
	p(now)

	// Crear una fecha específica
	then := time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Crear una fecha específica en la zona horaria local
	thenLocal := time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.Local)
	p(thenLocal)

	// Sumar una hora a una fecha
	then = then.Add(1 * time.Hour)
	p(then)

	// Sumar un día a una fecha
	then = then.Add(24 * time.Hour)
	p(then)

	// Sumar 7 días a una fecha
	then = then.Add(24 * time.Hour * 7)
	p(then)

	// Obtener componentes de la fecha
	p(then.Year())       // Año
	p(then.Month())      // Mes
	p(then.Day())        // Día
	p(then.Hour())       // Hora
	p(then.Second())     // Segundos
	p(then.Nanosecond()) // Nanosegundos
	p(then.Location())   // Zona horaria
	p(then.Weekday())    // Día de la semana

	// Obtener componentes de la fecha y hora actual
	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Second())
	p(now.Nanosecond())
	p(now.Location())
	p(now.Weekday())

	// Comparaciones entre fechas
	p("Then es antes que Now: ", then.Before(now))  // true
	p("Then es después que Now: ", then.After(now)) // false
	p("Then es igual que Now: ", then.Equal(now))   // false

	// Calcular diferencia entre fechas
	diff := now.Sub(then)
	p(diff)

	// Diferencia en horas entre dos fechas
	p("Diferencia en horas de now sobre then: ", diff.Hours())

	// Diferencia en días entre dos fechas
	p("Diferencia en días de now sobre then: ", diff.Hours()/24)

	// Diferencia en años entre dos fechas
	p("Diferencia en años de now sobre then: ", (diff.Hours()/24)/365)

	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Paquete OS: operaciones relacionadas con el sistema operativo
	fmt.Println(`
	---------------- Package  OS ----------------
	`)

	// Abrir un archivo
	file, err := os.Open("main.go")
	if err != nil {
		p(err)
		os.Exit(1) // Salir del programa si hay un error
	}
	defer file.Close() // Cerrar el archivo al finalizar

	p(file)

	// Obtener información sobre el archivo
	v, _ := file.Stat()
	p("Nombre del archivo: ", v.Name())
	p("El archivo pesa: ", v.Size(), " bytes")

	// Leer el archivo
	data := make([]byte, 4096) // Crear un búfer de lectura
	c, err := file.Read(data)  // Leer el contenido del archivo en el búfer
	if err != nil {
		p(err)
		os.Exit(1)
	}
	p(data, c)

	// Mostrar solo los datos leídos del archivo
	p(data[:c], c)

	// Mostrar los datos leídos del archivo como una cadena
	p(string(data[:c]), c)

	// Mostrar los datos leídos del archivo con formato
	fmt.Printf("read: %d bytes: %q\n", c, data[:c])

	// Obtener el valor de una variable de entorno
	p(os.Getenv("USERNAME"))

	// Modificar una variable de entorno
	os.Setenv("MI_ENV", "my value")
	p(os.Getenv("MI_ENV"))

	// Obtener el directorio actual de trabajo
	dir, _ := os.Getwd()
	p(dir)

	// Verificar si una variable de entorno existe
	p(os.Getenv("ENV_EXISTS"))
	p(os.Getenv("ENV_NOT_EXISTS"))

	// Buscar el valor de una variable de entorno
	env, ok := os.LookupEnv("ENV_EXISTS")
	p(env, ok)

	// Buscar el valor de una variable de entorno que no existe
	envFalse, ok := os.LookupEnv("ENV_NOT_EXISTS")
	p(envFalse, ok)

	// Ejemplo de expansión de variables de entorno en una URL de base de datos
	os.Setenv("DB_USERNAME", "nahuel")
	os.Setenv("DB_PASSWORD", "mysuperpassword")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "users")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT$DB_NAME")
	p(dbURL)

	fmt.Println(`
	---------------- Package  Log ----------------
	`)

}
