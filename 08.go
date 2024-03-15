package main

import "fmt"

func main08() {
	var metros float64
	fmt.Print("Digite uma medida: ")
	fmt.Scan(&metros)
	cent := metros / 100
	mili := metros / 1000
	fmt.Printf("%.2f em centímetros é %.2f e em milímetros é %.2f", metros, cent, mili)
}
