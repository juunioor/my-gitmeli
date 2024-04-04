package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Produto struct {
	id    int
	preco float64
	qtd   int
}

func guardarInfos(produtos []Produto) {
	var content strings.Builder
	for _, produto := range produtos {
		infos := fmt.Sprintf("%v, %.2f, %d\n", produto.id, produto.preco, produto.qtd)
		content.WriteString(infos)

		err := os.WriteFile("dados.csv", []byte(content.String()), 0644)
		if err != nil {
			fmt.Printf("[!] Erro ao escrever no arquivo: %v\n", err)
		} else {
			fmt.Printf("[+] As infos: %v    foram inseridas com sucesso!\n\n", infos)
		}
	}
}

// ..............................................................................................................................

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

// ..............................................................................................................................

func main() {
	produtos := []Produto{
		{id: 11, preco: 25.50, qtd: 1},
		{id: 12, preco: 10.15, qtd: 2},
		{id: 13, preco: 114.99, qtd: 2},
		{id: 14, preco: 24.69, qtd: 5},
	}

	guardarInfos(produtos)
	lerInfos("infos.csv")
}
