package main

import "fmt"

func analiseEmprestimo() {
	type Cliente struct {
		idade          int
		empregado      bool
		tempoDeEmprego int
		salario        float32
	}

	clientes := []Cliente{
		{idade: 21, empregado: true, tempoDeEmprego: 2, salario: 120000.00},
		{idade: 30, empregado: false, tempoDeEmprego: 3, salario: 80000.00},
		{idade: 44, empregado: true, tempoDeEmprego: 8, salario: 105000.00},
		{idade: 28, empregado: true, tempoDeEmprego: 4, salario: 110000.00},
	}

	for i := 0; i < len(clientes); i++ {
		cliente := clientes[i]
		if cliente.idade > 22 && cliente.empregado && cliente.tempoDeEmprego > 1 {
			fmt.Printf("[%v] Você está elegível para o empréstimo!\n", i)
			if cliente.salario > 100000 {
				fmt.Printf("[%v] Como você possui o salário maior que US$ 100.000, você está isento de juros!\n", i)
			} else {
				fmt.Printf("[%v] Como você possui o salário menor que US$ 100.000, você não está isento de juros.\n", i)
			}
		} else {
			fmt.Printf("[%v] Você não está elegível para o empréstimo.\n", i)
		}
	}
}

func main() {
	analiseEmprestimo()
	fmt.Println()
}
