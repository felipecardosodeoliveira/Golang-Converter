package main

import "fmt"

func main() {
	// Ponto de ebulição da água em Kelvin
	kelvin := 373.15

	// Conversão de Kelvin para Celsius
	celsius := kelvin - 273.15

	// Exibir o resultado
	fmt.Printf("O ponto de ebulição da água é %.2f °C\n", celsius)
}
