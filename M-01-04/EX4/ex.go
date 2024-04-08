package main

import "fmt"

func imprimirDados() {
	var sobrenome string = "Silva"

	// Para atribuir int a uma variável deve-se colocar apenas o número, e não com aspas.
	var idade int = 25 // var idade int = "25" > var idade int = 25

	// Para atribuir o booleano falso na variável deve-se usar sem aspas, pois assim está declarando uma string.
	boolean := false // boolean := "false" > boolean := false

	// Para salário e valores quebrados, usamos o float.
	var salario float32 = 4585.90 // var salario string = 4585.90 > var salario float32 = 4585
	var nome string = "Fellipe"

	fmt.Printf("%v %v possui %d anos. Sua habilitação consta como %v. Seu salário é R$ %.2f.\n", nome, sobrenome, idade, boolean, salario)
}

func main() {
	imprimirDados()
	fmt.Println()
}
