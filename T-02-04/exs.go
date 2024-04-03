package main

import "fmt"

type Aluno struct {
	nome      string
	sobrenome string
	rg        string
	data      string
}

func (aluno Aluno) Detalhamento() {
	fmt.Printf("Nome: [%v]\n", aluno.nome)
	fmt.Printf("Sobrenome: [%v]\n", aluno.sobrenome)
	fmt.Printf("RG: [%v]\n", aluno.rg)
	fmt.Printf("Data de admissão: [%v]\n", aluno.data)
}

// ..............................................................................................................................

type loja struct {
	produtos []IntefaceProduto
}

type produto struct {
	tipo  string
	nome  string
	preco float64
}

type IntefaceProduto interface {
	calcularCusto() float64
}

type InterfaceEcommerce interface {
	Total() float64
	Adicionar(produto IntefaceProduto)
}

func novoProduto(tipo, nome string, preco float64) IntefaceProduto {
	return produto{tipo: tipo, nome: nome, preco: preco}
}

func novaLoja() InterfaceEcommerce {
	return &loja{}
}

func (produto produto) calcularCusto() float64 {
	switch produto.tipo {
	case "Pequeno":
		return produto.preco

	case "Médio":
		return produto.preco * 1.03

	case "Grande":
		return produto.preco*1.06 + 2500

	default:
		return produto.preco
	}
}

func (loja loja) Total() float64 {
	total := 0.0
	for _, produto := range loja.produtos {
		total += produto.calcularCusto()
	}
	return total
}

func (loja *loja) Adicionar(produto IntefaceProduto) {
	loja.produtos = append(loja.produtos, produto)
}

func main() {
	aluno1 := Aluno{
		nome:      "Junior",
		sobrenome: "Lopes",
		rg:        "00.000.000-0",
		data:      "02/04/2024",
	}

	aluno1.Detalhamento()
	fmt.Printf("\n") // .........................................................................................................

	minhaLoja := novaLoja()

	produto1 := novoProduto("Pequeno", "Caneta", 2.50)
	produto2 := novoProduto("Médio", "Notebook", 1200.00)
	produto3 := novoProduto("Grande", "Sofá", 800.00)

	minhaLoja.Adicionar(produto1)
	minhaLoja.Adicionar(produto2)
	minhaLoja.Adicionar(produto3)

	fmt.Printf("Total da loja: R$%.2f\n", minhaLoja.Total())

}
