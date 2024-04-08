package main

import "fmt"

func imprimirNome() {
	var (
		nome  string = "Valdir Junior"
		idade int    = 22
	)
	fmt.Printf("%v tem %v anos. \n", nome, idade)
}

func main() {
	imprimirNome()
	fmt.Printf("\n")
}
