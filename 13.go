package main

import "fmt"

func main13() {
	var salario float32
	fmt.Print("Digite o valor do seu salário: ")
	fmt.Scan(&salario)
	aumento := salario + (salario * 15 / 100)
	fmt.Println("O seu salário antigo de R$", salario, "foi para R$", aumento, "com o aumento de 15%")
}
