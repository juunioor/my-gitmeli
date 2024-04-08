package main

import (
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

func main() {
	produtos := []Produto{
		{id: 11, preco: 25.50, qtd: 1},
		{id: 12, preco: 10.15, qtd: 2},
		{id: 13, preco: 114.99, qtd: 2},
		{id: 14, preco: 24.69, qtd: 5},
	}

	guardarInfos(produtos)
	fmt.Println()
}
