package main

import "fmt"

type Usuario struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

func mudarNome(usuario *Usuario, novoNome string) {
	usuario.nome = novoNome
}

func mudarIdade(usuario *Usuario, novaIdade int) {
	usuario.idade = novaIdade
}

func mudarEmail(usuario *Usuario, novoEmail string) {
	usuario.email = novoEmail
}

func mudarSenha(usuario *Usuario, novaSenha string) {
	usuario.senha = novaSenha
}

func main() {
	u1 := Usuario{
		nome:      "Lopes",
		sobrenome: "Junior",
		idade:     22,
		email:     "valdir.junior@msn.com",
		senha:     "12345",
	}

	mudarNome(&u1, "Valdir")
	mudarIdade(&u1, 22)
	mudarEmail(&u1, "valdir.junior@gmail.com")
	mudarSenha(&u1, "123")

	fmt.Println("[*]", u1)
}
