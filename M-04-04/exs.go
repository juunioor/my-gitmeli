package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "erro: O salário digitado não está dentro do valor mínimo"
}

type ErroHora struct{}

func (e *ErroHora) Error() string {
	return "erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês"
}

func salarioMensal(horas int, valor float64) (float64, error) {
	salarioMensal := valor * float64(horas)
	if horas < 80 {
		return 0, &ErroHora{}
	}
	if valor < 1 {
		return 0, fmt.Errorf("erro: o valor de horas é inválido (o valor é igual a 0 ou negativo)")
	}
	if salarioMensal >= 20000 {
		salarioMensal *= 0.9
	}

	return salarioMensal, nil
}

func main() {
	salario := 14999

	fmt.Println("EX1")
	if salario < 15000 {
		err := &MyError{}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
	fmt.Printf("\n") // .........................................................................................................

	fmt.Println("EX2")
	if salario < 15000 {
		err := errors.New("erro: O salário digitado não está dentro do valor mínimo")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
	fmt.Printf("\n") // .........................................................................................................

	fmt.Println("EX3")
	if salario < 15000 {
		err := fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
	fmt.Printf("\n") // .........................................................................................................

	horas := 90
	valor := -2.0

	fmt.Println("EX4")
	salarioMensal, err := salarioMensal(horas, valor)
	if err != nil {
		if errors.Unwrap(err) != nil {
			fmt.Println(errors.Unwrap(err).Error())
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Printf("O salário mensal é: %.2f\n", salarioMensal)
	}
}
