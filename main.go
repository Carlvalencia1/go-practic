package main 

/*
Declara una variable nombre que guarde tu nombre.
Declara una variable anio con el año actual.
Muestra ambos valores en consola.

import "fmt"

func main() {
	var name string= "Carlos Daniel Valencia Díaz"
	anio := 2025

	fmt.Println("Mi nombre:", name, "Año:", anio)

} 

import (
	"fmt"
	"time"
)

func main() {
	var name string
	var pais string
	var edad int 

	fmt.Println("¿Cuál es tu nombre?")
	fmt.Scanln(&name)
	fmt.Println("¿Cuál es tu país?")
	fmt.Scanln(&pais)
	fmt.Println("¿Cuál es tu edad?")
	fmt.Scanln(&edad)

	anioActual := time.Now().Year()
	anioNacimiento := anioActual - edad

	fmt.Printf("Hola %s de %s, naciste en %d\n", name, pais, anioNacimiento)
}
	*/

	import "fmt"

func main() {
    edad := 18

    if edad >= 18 {
        fmt.Println("Eres mayor de edad.")
    } else {
        fmt.Println("Eres menor de edad.")
    }
}

//add estructuras de control