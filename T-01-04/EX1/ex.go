package main

import "fmt"

func imprimirLetras(palavra string) {
	fmt.Printf("A quantidade de letras que a palavra %q tem é: %v.\n", palavra, len(palavra))
	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%c\n", palavra[i])
	}
}

func main() {
	imprimirLetras("banana")
	fmt.Println()
}
