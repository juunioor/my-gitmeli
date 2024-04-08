package main

import "fmt"

type Product struct {
	nome  string
	preco float64
	qtd   int
}

type Servicos struct {
	nome    string
	preco   float64
	minutos int
}

type Manutencao struct {
	nome  string
	preco float64
}

func somarProdutos(produtos []Product, result chan<- float64) {
	precototal := 0.00
	for _, produto := range produtos {
		precototal += produto.preco * float64(produto.qtd)
	}
	result <- precototal
}

func SomarServicos(servicos []Servicos, result chan<- float64) {
	precototal := 0.00
	minutos := 0
	for _, servico := range servicos {
		precototal += servico.preco
		minutos += servico.minutos
	}
	mediaHoras := (minutos / len(servicos))
	if mediaHoras < 30 {
		precototal *= 30
	} else {
		precototal *= float64(mediaHoras)
	}

	result <- precototal
}

func somarManutencao(manutencoes []Manutencao, result chan<- float64) {
	precototal := 0.00
	for _, manutencao := range manutencoes {
		precototal += manutencao.preco
	}
	result <- precototal
}

func main() {
	produtos := []Product{
		{"Camiseta", 10.00, 2},
		{"Calça", 20.00, 1},
		{"Meia", 5.50, 3},
	}
	servicos := []Servicos{
		{"Lavagem", 50.00, 60},
		{"Reparação", 30.00, 45},
		{"Costura", 40.00, 90},
	}
	manutencoes := []Manutencao{
		{"Coloração", 100.00},
		{"Golas", 150.00},
		{"Alargamento", 200.00},
	}

	resultProdutos := make(chan float64)
	resultServicos := make(chan float64)
	resultManutencoes := make(chan float64)

	go somarProdutos(produtos, resultProdutos)
	go SomarServicos(servicos, resultServicos)
	go somarManutencao(manutencoes, resultManutencoes)

	totalProdutos := <-resultProdutos
	totalServicos := <-resultServicos
	totalManutencoes := <-resultManutencoes

	fmt.Printf("[!] Total final: %.2f\n", totalProdutos+totalServicos+totalManutencoes)
}
