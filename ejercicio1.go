package main

import "fmt"

func main() {
	var nombre string
	var direccion string

	fmt.Println("Ingrese su nombre")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su direccion")
	fmt.Scanln(&direccion)

	fmt.Println(nombre, direccion)

}
