package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Produto struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	Cor         string    `json:"cor"`
	Preco       float64   `json:"preco"`
	Estoque     int       `json:"estoque"`
	Codigo      string    `json:"codigo"`
	Publicacao  bool      `json:"publicacao"`
	DataCriacao time.Time `json:"data_criacao"`
}

func main() {
	produto := Produto{
		ID:          1,
		Nome:        "Camiseta",
		Cor:         "Azul",
		Preco:       25.99,
		Estoque:     100,
		Codigo:      "ABC123",
		Publicacao:  true,
		DataCriacao: time.Now(),
	}

	infos, err := json.Marshal(produto)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(infos))
}
