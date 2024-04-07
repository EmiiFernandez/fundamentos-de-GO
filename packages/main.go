package main

import (
	"fmt"
	"log" // Importar el paquete log para el manejo de logs
	"os"
	"time"
)

func main() {

	p := fmt.Println

	// Paquete Time: operaciones relacionadas con el tiempo
	fmt.Println(`
	---------------- Package  Time ----------------
	`)

	now := time.Now() // Fecha y hora actual
	p(now)

	// Crear una fecha específica
	then := time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Crear una fecha específica en la zona horaria local
	thenLocal := time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.Local)
	p(thenLocal)

	// Agregar una hora a una fecha
	then = then.Add(1 * time.Hour)
	p(then)

	// Agregar un día a una fecha
	then = then.Add(24 * time.Hour)
	p(then)

	// Agregar 7 días a una fecha
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

	// Diferencias entre fechas
	diff := now.Sub(then) // Hacer resta entre fechas
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

	// Utiliza las funcionalidades del sistema operativo. Por ej: abrir un archivo

	file, err := os.Open("main.go")

	if err != nil {
		p(err)
		os.Exit(1) // Finaliza el programa. Con 1 --> ha ocurrido un error
	}
	p(file)

	v, _ := file.Stat() // Obtener información del archivo
	p("Nombre del archivo: ", v.Name())
	p("El archivo pesa: ", v.Size(), " bytes")

	// Leer el archivo
	data := make([]byte, 4096) // Crear cadena de bytes

	// Cadena de caracteres
	c, err := file.Read(data) // Almacenar la información de file
	if err != nil {
		p(err)
		os.Exit(1)
	}
	p(data, c)

	p(data[:c], c)

	// Mostrar lo que contiene el archivo
	p(string(data[:c]), c)

	// %q convierte los bytes a formato tipo texto
	fmt.Printf("read: %d bytes: %q\n", c, data[:c])

	p(os.Getenv("USERNAME")) // Devuelve el valor de una determinada variable de entorno. Si la variable no existe, devuelve vacío

	p(os.Getenv("MI_ENV"))
	os.Setenv("MI_ENV", "my value") // una variable de entorno
	p(os.Getenv("MI_ENV"))

	dir, _ := os.Getwd() // Imprime la dirección donde se encuentra
	p(dir)

	os.Setenv("ENV_EXISTS", "")
	p(os.Getenv("ENV_EXISTS"))
	p(os.Getenv("ENV_NOT_EXISTS"))

	// LookupEnv --> Saber si existe una variable de entorno
	env, ok := os.LookupEnv("ENV_EXISTS")
	p(env, ok) // Imprime primero vacío (porque es la variable que tiene)

	envFalse, ok := os.LookupEnv("ENV_NOT_EXISTS")
	p(envFalse, ok)

	os.Setenv("DB_USERNAME", "nahuel")
	os.Setenv("DB_PASSWORD", "mysuperpassword")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "users")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT$DB_NAME")
	p(dbURL)

	// Imprimir un encabezado indicando el uso del paquete Log
	fmt.Println(`
		---------------- Package  Log ----------------
		`)

	// Imprimir un mensaje en el registro de logs
	log.Println("test")

	// Crear un archivo de registro con el timestamp actual como nombre
	date := time.Now()
	fileLog, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano())) // Creo archivo de log
	if err != nil {
		log.Panic(err.Error()) // Salir del programa si hay un error al crear el archivo de log
	}

	// Crear un nuevo logger que escriba en el archivo de log con prefijo "success: " y flags de fecha y archivo cortos
	l := log.New(fileLog, "success: ", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")

	// Crear un nuevo logger que escriba en el mismo archivo de log con prefijo "success: " y solo flags de fecha
	l2 := log.New(fileLog, "success: ", log.LstdFlags)
	l2.Println("test 1")
	l2.Println("test 2")
	l2.Println("test 3")

	// Crear un nuevo logger que escriba en el mismo archivo de log sin prefijo y sin flags
	l3 := log.New(fileLog, "success: ", 0)
	l3.Println("test 1")
	l3.Println("test 2")
	l3.Println("test 3")

	// Cambiar el prefijo del logger l a "errors: "
	l.SetPrefix("errors: ")
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")

	// Crear un nuevo logger que escriba en la salida estándar (consola) con prefijo "success: " y flags de fecha y archivo cortos
	l4 := log.New(os.Stdout, "success: ", log.LstdFlags|log.Lshortfile)
	l4.Println("test 1")
	l4.Println("test 2")
	l4.Println("test 3")
}
