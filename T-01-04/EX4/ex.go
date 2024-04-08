package main

import "fmt"

func imprimirDados() {
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
	imprimirDados()
	fmt.Println()
}
