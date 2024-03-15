package main

import "fmt"

func main05() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	suc := num + 1
	ant := num - 1
	fmt.Printf("O seu antecessor é %d e sucessor é %d", ant, suc)
}
