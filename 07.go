package main

import "fmt"

func main07() {
	var nota1 float64
	var nota2 float64
	fmt.Print("Informe duas notas: ")
	fmt.Scan(&nota1, &nota2)
	media := (nota1 + nota2) / 2
	fmt.Printf("Sua média é %.2f", media)
}
