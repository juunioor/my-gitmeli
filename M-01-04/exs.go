package main

import "fmt"

func Exercicio1() {
	var (
		nome  string = "Valdir Junior"
		idade int    = 22
	)
	fmt.Printf("%v tem %v anos. \n", nome, idade)
}

func Exercicio2() {
	var (
		temperatura        int8    = 29
		umidade            float32 = 39.2
		pressaoAtmosferica int16   = 1019
	)

	fmt.Printf("A temperatura hoje está é %v °C, com %v para umidade. A pressão atmosférica é %vhPa. \n", temperatura, umidade, pressaoAtmosferica)
}

func Exercicio3() {
	// Não se pode utilizar números como primeiro dígito na declaração de variáveis.
	var nome string = "Junior" // var 1nome string > var nome string
	var sobrenome string = "Lopes"

	// O tipo da variável deve vir após o nome da variável.
	var idade int = 22 // var int idade > var idade int

	// Seguindo na mesma ideia, não se pode declarar variáveis com números na frente.
	// Ainda sim, não é possível sobreescrever um valor já existente em uma variável, assim mudaremos o nome. É recomendável utilizar nome de variáveis claras.
	sobrenome1 := 6 // 1sobrenome := 6 > sobrenome1 := 6
	var licenca_para_dirigir = true

	// Para declaração de variáveis, não pode usar-se espaço, deve-se usar camelCase. A escolha da variável deve estar de acordo com o valor que irá receber.
	var estaturaDaPessoa float32 = 1.74 // var estatura da pessoa int > var estaturaDaPessoa float32
	quantidadeDeFilhos := 2

	fmt.Printf("%s %s tem %d, possui %v sobrenomes em sua família, sua habilitação consta como %v, possui %v de altura e possui %d filhos.\n", nome, sobrenome, idade, sobrenome1, licenca_para_dirigir, estaturaDaPessoa, quantidadeDeFilhos)
}

func Exercicio4() {
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
	Exercicio1()
	fmt.Println("\n")

	Exercicio2()
	fmt.Println("\n")

	Exercicio3()
	fmt.Println("\n")

	Exercicio4()
	fmt.Println("\n")
}
