package main

import (
	"errors"
	"fmt"
)

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

func calcularEstatistica(operation string) (func(...int) float64, error) {
	switch operation {
	case "minimum":
		return opMin, nil

	case "average":
		return opAvg, nil

	case "maximum":
		return opMax, nil

	default:
		return nil, errors.New("operação inválida")
	}
}

func main() {
	minhaFunc, err := calcularEstatistica(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}

	averageFunc, err := calcularEstatistica(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxFunc, err := calcularEstatistica(maximum)
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
}
