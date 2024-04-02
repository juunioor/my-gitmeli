package main

import (
	"fmt"
)

func Exercicio1() {
	palavra := "banana"
	fmt.Printf("A quantidade de letras que a palavra %q tem é: %v.\n", palavra, len(palavra))
	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%c\n", palavra[i])
	}
}

func Exercicio2() {
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
		if cliente.idade > 22 && cliente.empregado == true && cliente.tempoDeEmprego > 1 {
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

func Exercicio3(mes int) string {
	/* Escolhi realizar esta solução pois temos uma determinada ordem de "equivalência" entre os meses,
	o que é possível assim, fazer com que o valor dado corresponda com o index equivalente.*/
	meses := [12]string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}

	if mes < 1 || mes > 12 {
		return "Valor inválido!"
	} else {
		return meses[mes-1]
	}
}

func Exercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("A idade de Benjamin é %v anos.\n", employees["Benjamin"])
	employees["Valdir"] = 22
	delete(employees, "Pedro")

	var count int
	for _, idade := range employees {
		if idade > 21 {
			count++
		}
	}
	fmt.Printf("Na lista possui %v pessoas acima de 21 anos.\n", count)

}

func main() {
	Exercicio1()
	fmt.Printf("\n")
	Exercicio2()
	fmt.Printf("\n")
	fmt.Println(Exercicio3(3))
	fmt.Printf("\n")
	Exercicio4()
}
