package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	arquivo   int
	nome      string
	sobrenome string
	rg        string
	telefone  string
	endereco  string
}

func lerDados(arquivo string) ([]Cliente, error) {
	file, err := os.Open(arquivo)
	if err != nil {
		panic(fmt.Errorf("o arquivo indicado não foi encontrado ou está danificado"))
	}
	defer file.Close()

	var clientes []Cliente
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linha := scanner.Text()
		pedacos := strings.Split(linha, ",")
		if len(pedacos) != 6 {
			return nil, fmt.Errorf("erro: o arquivo não contém todos os dados necessários")
		}

		arquivo, err := strconv.Atoi(pedacos[0])
		if err != nil {
			return nil, fmt.Errorf("erro: o ID está em formato errado")
		}

		cliente := Cliente{
			arquivo:   arquivo,
			nome:      pedacos[1],
			sobrenome: pedacos[2],
			rg:        pedacos[3],
			telefone:  pedacos[4],
			endereco:  pedacos[5],
		}
		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func adicionarCliente(arquivo string, clientes []Cliente) error {
	file, err := os.Create(arquivo)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo")
	}
	defer file.Close()

	for _, cliente := range clientes {
		_, err := fmt.Fprintf(file, "%v,%v,%v,%v,%v,%v\n", cliente.arquivo, cliente.nome, cliente.sobrenome, cliente.rg, cliente.telefone, cliente.endereco)
		if err != nil {
			return fmt.Errorf("erro ao escrever no arquivo")
		}
	}
	return nil
}

func verificarCliente(clientes []Cliente, cliente Cliente) error {
	for _, c := range clientes {
		if c.rg == cliente.rg {
			return fmt.Errorf("cliente já existe")
		}
	}
	return nil
}

func gerarID(arquivo string) (int, error) {
	clientes, err := lerDados(arquivo)
	if err != nil {
		return 0, err
	}
	return len(clientes) + 1, nil
}

func main() {
	arquivo := "customers.txt"

	// criando ID do usuário
	id, err := gerarID(arquivo)
	if err != nil {
		panic(err)
	}

	// criando clienteNovo
	clienteNovo := Cliente{
		arquivo:   id,
		nome:      "Valdir",
		sobrenome: "Junior",
		rg:        "111.111.111-11",
		telefone:  "(11) 11111-1111",
		endereco:  "Rua dos desconhecidos",
	}

	dados, err := lerDados(arquivo)
	if err != nil {
		panic(err)
	}

	err = verificarCliente(dados, clienteNovo)
	if err != nil {
		panic(err)
	}

	clientes := append(dados, clienteNovo)
	fmt.Println(clientes)

	adicionarCliente(arquivo, clientes)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()
}
