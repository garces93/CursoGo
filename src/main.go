package main

import "fmt"

func main() {
	//Declaracion de Constantes que jamas va a cambiar con el tiempo
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println(pi, pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area Cuadrado
	const BaseCuadrado = 10
	areaCuadrado := BaseCuadrado * BaseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

}
