package main

import "fmt"

func calcularImposto(salario float64) {
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

func main() {
	calcularImposto(8000.20)
	fmt.Println()

}
