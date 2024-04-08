package main

import "fmt"

func imprimirVariaveis() {
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

func main() {
	imprimirVariaveis()
	fmt.Println()
}
