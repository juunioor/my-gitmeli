package main

import (
	"fmt"
	"os"
)

func lerArquivo(arquivo string) {
	file, err := os.Open(arquivo)
	if err != nil {
		panic(fmt.Errorf("o arquivo indicado não foi encontrado ou está danificado"))
	}
	defer file.Close()
}

func main() {
	arquivo := "customers.txt"
	lerArquivo(arquivo)
	fmt.Println()
}
