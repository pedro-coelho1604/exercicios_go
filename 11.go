package main

import (
	"fmt"
)

func main11() {
	var largura, altura float64
	fmt.Println("Digite a largura e altura da parede em metros:")
	fmt.Scanln(&largura, &altura)
	area := largura * altura
	quantidadeTinta := area / 2
	fmt.Printf("Serão necessários %.2f litros de tinta para pintar sua parede", quantidadeTinta)
}
