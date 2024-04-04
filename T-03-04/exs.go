package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Usuario struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

func mudarNome(usuario *Usuario, novoNome string) {
	usuario.nome = novoNome
}

func mudarIdade(usuario *Usuario, novaIdade int) {
	usuario.idade = novaIdade
}

func mudarEmail(usuario *Usuario, novoEmail string) {
	usuario.email = novoEmail
}

func mudarSenha(usuario *Usuario, novaSenha string) {
	usuario.senha = novaSenha
}

// ..............................................................................................................................

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

// ..............................................................................................................................

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

// ..............................................................................................................................

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	u1 := Usuario{
		nome:      "Lopes",
		sobrenome: "Junior",
		idade:     22,
		email:     "valdir.junior@msn.com",
		senha:     "12345",
	}

	mudarNome(&u1, "Valdir")
	mudarIdade(&u1, 22)
	mudarEmail(&u1, "valdir.junior@gmail.com")
	mudarSenha(&u1, "123")

	fmt.Println("[EX1]")
	fmt.Println("[*]", u1)
	fmt.Println() // .............................................................................................................

	cli1 := Cliente{
		nome:      "Rapha",
		sobrenome: "Gumieri",
		email:     "rapha.gumieri@gmail.com",
		produtos:  []Produto{},
	}

	fmt.Println("[EX2]")
	fmt.Printf("[+] Produto inserido: %v\n", novoProduto("Iphone 11", 2450.99))

	adicionarProduto(&cli1, novoProduto("Iphone 11", 2450.99), 2)
	adicionarProduto(&cli1, novoProduto("Iphone 11", 2450.99), 5)

	fmt.Printf("[!] Os produtos do cliente %v são %v.\n", cli1.nome, cli1.produtos)

	deletarProdutos(&cli1)

	fmt.Printf("[!] Os produtos do cliente %v são %v.\n", cli1.nome, cli1.produtos)
	fmt.Println() // .............................................................................................................

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

	fmt.Println("[EX3]")
	fmt.Printf("[!] Total final: %.2f\n", totalProdutos+totalServicos+totalManutencoes)
	fmt.Println() // .............................................................................................................

	array100 := rand.Perm(100)
	array1000 := rand.Perm(1000)
	array10000 := rand.Perm(10000)

	fmt.Println("[EX4]")

	// insertionSort
	start := time.Now()
	insertionSort(array100)
	fmt.Println("[+] Tempo para 100 elementos usando insertionSort:", time.Since(start))

	start = time.Now()
	insertionSort(array1000)
	fmt.Println("[+] Tempo para 1000 elementos usando insertionSort:", time.Since(start))

	start = time.Now()
	insertionSort(array10000)
	fmt.Println("[+] Tempo para 10000 elementos usando insertionSort:]", time.Since(start))
	fmt.Println()

	// bubbleSort
	start = time.Now()
	bubbleSort(array100)
	fmt.Println("[+] Tempo para 100 elementos usando bubbleSort:", time.Since(start))

	start = time.Now()
	bubbleSort(array1000)
	fmt.Println("[+] Tempo para 1000 elementos usando bubbleSort:", time.Since(start))

	start = time.Now()
	bubbleSort(array10000)
	fmt.Println("[+] Tempo para 10000 elementos usando bubbleSort:", time.Since(start))
	fmt.Println()

	// selectionSort
	start = time.Now()
	selectionSort(array100)
	fmt.Println("[+] Tempo para 100 elementos usando selectionSort:", time.Since(start))

	start = time.Now()
	selectionSort(array1000)
	fmt.Println("[+] Tempo para 1000 elementos usando selectionSort:", time.Since(start))

	start = time.Now()
	selectionSort(array10000)
	fmt.Println("[+] Tempo para 10000 elementos usando selectionSort:]", time.Since(start))
	fmt.Println() // .............................................................................................................
}
