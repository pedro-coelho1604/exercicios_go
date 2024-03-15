package main

import "fmt"

func main02() {
	var nome string
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)
	fmt.Printf("Boas vindas %s", nome)
}
