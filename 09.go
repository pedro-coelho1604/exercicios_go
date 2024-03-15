package main

import "fmt"

func main09() {
	var numero int
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)
	for i := 0; i <= 10; i++ {
		mult := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, mult)
	}
}
