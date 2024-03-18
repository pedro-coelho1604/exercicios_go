package main

import (
	"fmt"
)

func main15() {
	var km float32
	var dias int
	fmt.Print("Quantos Km você rodou, e quantos dias você está com o carro ? ")
	fmt.Scan(&km, &dias)
	precoKm := km * 0.15
	precodia := dias * 60
	cobranca := precoKm + float32(precodia)
	fmt.Println("Você irá pagar", cobranca, "pelo aluguel do carro")
}
