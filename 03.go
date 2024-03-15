package main

import "fmt"

func main03() {
	var num1 int
	var num2 int
	fmt.Print("Digite dois números: ")
	fmt.Scan(&num1, &num2)
	soma := num1 + num2
	fmt.Printf("A soma dos dois números é %d", soma)
}
