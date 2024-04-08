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

func main() {
	aluno1 := Aluno{
		nome:      "Junior",
		sobrenome: "Lopes",
		rg:        "00.000.000-0",
		data:      "02/04/2024",
	}

	aluno1.Detalhamento()
	fmt.Println()
}
