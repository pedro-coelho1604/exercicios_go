package main

import (
	"fmt"
	"math"
)

func main06() {
	var num float64
	fmt.Print("Digite seu número: ")
	fmt.Scan(&num)
	dobro := num * 2
	triplo := num * 3
	raiz := math.Sqrt(num)
	fmt.Printf("O dobro é %.2f, o triplo é %.2f e a raiz é %.2f", dobro, triplo, raiz)
}
