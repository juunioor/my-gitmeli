package main

import (
	"errors"
	"fmt"
)

func Exercicio1(salario float64) {
	if salario > 50000 && salario < 150000 {
		imposto := salario * 0.17
		fmt.Printf("Foi descontado R$%.2f de imposto do seu salário. Seu salário atual é de R$%.2f.\n", imposto, salario-imposto)
	}
	if salario > 150000 {
		imposto := salario * 0.27
		fmt.Printf("Foi descontado R$%.2f de imposto do seu salário. Seu salário atual é de R$%.2f.\n", imposto, salario-imposto)
	} else {
		fmt.Println("Não foi descontado imposto do seu sálario!")
	}
}

// ..............................................................................................................................

func Exercicio2(notas ...float64) {
	var soma float64

	for _, nota := range notas {
		soma += nota
	}

	fmt.Printf("Sua média é %.2f. \n", soma/float64(len(notas)))
}

// ..............................................................................................................................

func Exercicio3() {
	var (
		categoria string
		horas     int
	)

	fmt.Printf("Digite a quantidade de horas trabalhadas: ")
	fmt.Scanln(&horas)

	fmt.Printf("Digite a categoria do funcionário: ")
	fmt.Scanln(&categoria)

	switch categoria {
	case "C":
		fmt.Printf("Seu salário é de R$%v.\n", horas*1000)

	case "B":
		if horas > 160 {
			fmt.Printf("Seu salário é de R$%v.\n", float64(horas*1500)*1.2)
		} else {
			fmt.Printf("Seu salário é de R$%v.\n", horas*1500)
		}

	case "A":
		if horas > 160 {
			fmt.Printf("Seu salário é de R$%v.\n", float64(horas*3000)*1.5)
		} else {
			fmt.Printf("Seu salário é de R$%v.\n", horas*3000)
		}

	default:
		fmt.Println("Categoria inválida.")
	}
}

// ..............................................................................................................................

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func opMin(valores ...int) float64 {
	min := 10

	for _, num := range valores {
		if num < min {
			min = num
		}
	}
	return float64(min)
}

func opAvg(valores ...int) float64 {
	var soma int

	for _, num := range valores {
		soma += num
	}
	return float64(soma / len(valores))
}

func opMax(valores ...int) float64 {
	max := 0

	for _, num := range valores {
		if num > max {
			max = num
		}
	}
	return float64(max)
}

func Exercicio4(operation string) (func(...int) float64, error) {
	switch operation {
	case "minimum":
		return opMin, nil

	case "average":
		return opAvg, nil

	case "maximum":
		return opMax, nil

	default:
		return nil, errors.New("Operação inválida!")
	}
}

// ..............................................................................................................................

const (
	cao       = "cao"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func opCao() float64 {
	var qtd int
	fmt.Printf("Digite a quantidade de Cães: ")
	fmt.Scanln(&qtd)
	return float64(qtd * 10)
}

func opGato() float64 {
	var qtd int
	fmt.Printf("Digite a quantidade de Gatos: ")
	fmt.Scanln(&qtd)
	return float64(qtd * 5)
}

func opHamster() float64 {
	var qtd int
	fmt.Printf("Digite a quantidade de Hamsters: ")
	fmt.Scanln(&qtd)
	return float64(qtd) * 0.25
}

func opTarantula() float64 {
	var qtd int
	fmt.Printf("Digite a quantidade de Tarântulas: ")
	fmt.Scanln(&qtd)
	return float64(qtd) * 0.15
}

func Exercicio5(animal string) (func() float64, error) {
	switch animal {
	case "cao":
		return opCao, nil

	case "gato":
		return opGato, nil

	case "hamster":
		return opHamster, nil

	case "tarantula":
		return opTarantula, nil

	default:
		return nil, errors.New("Animal inválido!")
	}
}

// ..............................................................................................................................

func main() {
	Exercicio1(8000.20)
	fmt.Printf("\n") //..........................................................................................................

	Exercicio2(9.5, 10, 2)
	fmt.Printf("\n") //..........................................................................................................

	Exercicio3()
	fmt.Printf("\n") //..........................................................................................................

	minhaFunc, err := Exercicio4(minimum)
	averageFunc, err := Exercicio4(average)
	maxFunc, err := Exercicio4(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimum:", minValue)
	fmt.Println("Average:", averageValue)
	fmt.Println("Maximum:", maxValue)
	fmt.Printf("\n") //..........................................................................................................

	caoFunc, err := Exercicio5(cao)
	gatoFunc, err := Exercicio5(gato)
	hamsterFunc, err := Exercicio5(hamster)
	tarantulaFunc, err := Exercicio5(tarantula)
	if err != nil {
		fmt.Println(err)
		return
	}

	caoValue := caoFunc()
	gatoValue := gatoFunc()
	hamsterValue := hamsterFunc()
	tarantulaValue := tarantulaFunc()

	fmt.Printf("\nCães: %v Kg. \n", caoValue)
	fmt.Printf("Gatos: %v Kg. \n", gatoValue)
	fmt.Printf("Hamsters: %v Kg. \n", hamsterValue)
	fmt.Printf("Tarântulas: %v Kg. \n", tarantulaValue)

}
