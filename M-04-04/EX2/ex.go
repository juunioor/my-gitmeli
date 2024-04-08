package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "erro: O salário digitado não está dentro do valor mínimo"
}

func main() {
	salario := 14999

	if salario < 15000 {
		err := errors.New("erro: O salário digitado não está dentro do valor mínimo")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
