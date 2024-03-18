package main

import "fmt"

func main16() {
	var num float32
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	numReal := int(num)
	fmt.Println("Seu número é:", numReal)
}
