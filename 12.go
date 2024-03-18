package main

import "fmt"

func main12() {
	var produto float32
	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&produto)
	produtoDesconto := produto - (produto * 5 / 100)
	fmt.Printf("O preço do produto com desconto fica por %.2f", produtoDesconto)
}
