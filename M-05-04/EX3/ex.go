package main

import (
	"time"

	"github.com/gin-gonic/gin"
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

var produtos = []Produto{
	{
		ID:          1,
		Nome:        "Camiseta",
		Cor:         "Azul",
		Preco:       25.99,
		Estoque:     100,
		Codigo:      "ABC123",
		Publicacao:  true,
		DataCriacao: time.Now(),
	},
	{
		ID:          1,
		Nome:        "Camiseta",
		Cor:         "Azul",
		Preco:       25.99,
		Estoque:     100,
		Codigo:      "ABC123",
		Publicacao:  true,
		DataCriacao: time.Now(),
	},
}

func GetAll(c *gin.Context) {
	c.JSON(200, produtos)
}

func HandleRequest() {
	r := gin.Default()
	r.GET("/products", GetAll)
	r.Run(":8080")
}

func main() {
	HandleRequest()
}
