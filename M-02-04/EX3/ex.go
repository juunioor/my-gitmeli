package main

import "fmt"

func calcularSalario() {
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

func main() {
	calcularSalario()
	fmt.Println()
}
