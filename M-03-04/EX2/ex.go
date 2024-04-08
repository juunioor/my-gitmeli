package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func lerInfos(arquivo string) {
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo!", err)
	}

	reader := csv.NewReader(file)

	linhas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo!", err)
	}

	for _, linha := range linhas {
		for _, coluna := range linha {
			fmt.Printf("%-20s", coluna)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	lerInfos("infos.csv")
	fmt.Println()
}
