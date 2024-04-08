package main

import "fmt"

func calcularMedia(notas ...float64) {
	var soma float64

	for _, nota := range notas {
		soma += nota
	}

	fmt.Printf("Sua média é %.2f. \n", soma/float64(len(notas)))
}

func main() {
	calcularMedia(9.5, 10, 2)
	fmt.Println()
}
