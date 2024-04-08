package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês"
}

func salarioMensal(horas int, valor float64) (float64, error) {
	salarioMensal := valor * float64(horas)
	if horas < 80 {
		return 0, &MyError{}
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
	horas := 90
	valor := -2.0

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
