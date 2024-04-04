package main

// Recordar para compilar se tiene que ocupar "go run FuncionesUtiles.go"

import (
	"fmt" //Formatting string y mostrando mensajes
	"strings"
)

func main() {

	var nombre string = "Nicolas"

	// Sprintf guarda una string formateada
	var nuevonombre string = fmt.Sprintf("Esto es un test %v", nombre)
	fmt.Println(nuevonombre)

	// Uso de arreglos en GO
	var array = [3]int{20, 25, 30}

	fmt.Println(array, len(array))

	// Uso de slices en GO  -> Son arreglos pero de otros lenguajes, donde su tamano puede crecer,
	// Se pueden agregar cosas y demaces. (Siguen siendo implementaciones de arreglos)
	var puntajes = []int{100, 200, 3000}
	puntajes = append(puntajes, 400)

	fmt.Println(puntajes, len(puntajes))

	// Tambien existen slices pero con rangos, para el caso incluye el numero menor pero excluye el ultimo
	rango1 := puntajes[1:3]
	rango2 := puntajes[2:]
	fmt.Println(rango1, rango2)

	//Ahora veremos el uso del paquete strings como ejemplo

	// Contains es una expresion regular
	bienvenida := "Hola que pasa mi gente"
	fmt.Println(strings.Contains(bienvenida, "Hol"))

	// ReplaceAll reemplaza la palabra deseada, mas no modifica el string original
	fmt.Println(strings.ReplaceAll(bienvenida, "Hola", "Olis"))

	// Index retorna la posicion dentro del string donde se encontro la palabra deseada
	fmt.Println(strings.Index(bienvenida, "que"))

	// Separa la string en un arreglo donde esta separado por el valor dado
	fmt.Println(strings.Split(bienvenida, " "))

	// AHORA VEMOS LA SINTAXSIS DE LOS CICLOS

	// Sintaxis para un ciclo loop normal
	for i := 0; i < 5; i++ {
		fmt.Println("El valor de i es: ", i)
	}

	// Sintaxis para un ciclo loop en un array/slice
	nombres := []string{"Mario", "Luigi", "Peach", "Daisy"}

	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	// Ciclo for corte Python
	for index, value := range nombres {
		fmt.Printf("La posicion en el indice %v es %v \n", index, value)
	}

	// Ciclo for corte Python pero mas cuentiado
	for _, valor := range nombres {
		fmt.Printf("El nombre es: %v \n", valor)
	}
}
