package main

import "fmt"

func main14() {
	var celsius float32
	fmt.Print("Digite sua temperatura em Cº: ")
	fmt.Scan(&celsius)
	fahr := ((9 * celsius) / 5) + 32
	fmt.Println("Sua temperatura", celsius, "em celsius corresponde a", fahr, "Fahrenheit")
}
