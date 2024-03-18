package main

import "fmt"

func main10() {
	var money float64
	fmt.Print("Quanto você tem na conta: ")
	fmt.Scan(&money)
	convert := money / 3.27
	fmt.Printf("Você pode comprar %.2f dólares", convert)
}
