package main

import "fmt"

type Cliente struct {
	nome      string
	sobrenome string
	email     string
	produtos  []Produto
}

type Produto struct {
	nome  string
	preco float64
	qtd   int
}

func novoProduto(nome string, preco float64) Produto {
	produto := Produto{
		nome:  nome,
		preco: preco,
	}
	return produto
}

func adicionarProduto(cliente *Cliente, novoProduto Produto, quantidade int) {
	produtoExistente := false

	for i, produto := range cliente.produtos {
		if produto.nome == novoProduto.nome {
			cliente.produtos[i].qtd += quantidade
			produtoExistente = true
			break
		}
	}
	if !produtoExistente {
		novoProduto.qtd = quantidade
		cliente.produtos = append(cliente.produtos, novoProduto)
	}
}

func deletarProdutos(cliente *Cliente) {
	cliente.produtos = []Produto{}
}

func main() {
	cli1 := Cliente{
		nome:      "Rapha",
		sobrenome: "Gumieri",
		email:     "rapha.gumieri@gmail.com",
		produtos:  []Produto{},
	}

	fmt.Printf("[+] Produto inserido: %v\n", novoProduto("Iphone 11", 2450.99))

	adicionarProduto(&cli1, novoProduto("Iphone 11", 2450.99), 2)
	adicionarProduto(&cli1, novoProduto("Iphone 11", 2450.99), 5)

	fmt.Printf("[!] Os produtos do cliente %v são %v.\n", cli1.nome, cli1.produtos)

	deletarProdutos(&cli1)

	fmt.Printf("[!] Os produtos do cliente %v são %v.\n", cli1.nome, cli1.produtos)
}
