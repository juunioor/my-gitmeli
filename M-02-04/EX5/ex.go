package main

import (
	"errors"
	"fmt"
)

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

func calcularRacao(animal string) (func() float64, error) {
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
		return nil, errors.New("animal inválido")
	}
}

func main() {
	caoFunc, err := calcularRacao(cao)
	if err != nil {
		fmt.Println(err)
		return
	}

	gatoFunc, err := calcularRacao(gato)
	if err != nil {
		fmt.Println(err)
		return
	}

	hamsterFunc, err := calcularRacao(hamster)
	if err != nil {
		fmt.Println(err)
		return
	}

	tarantulaFunc, err := calcularRacao(tarantula)
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
