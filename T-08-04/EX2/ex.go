package main

import (
	"strconv"
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
		ID:          2,
		Nome:        "Sapato",
		Cor:         "Vermelho",
		Preco:       67.99,
		Estoque:     56,
		Codigo:      "DFG123",
		Publicacao:  true,
		DataCriacao: time.Now(),
	},
}

func GetAll(c *gin.Context) {
	c.JSON(200, produtos)
}

func SearchProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "ID inválido.")
		return
	}

	for _, produto := range produtos {
		if produto.ID == id {
			c.JSON(200, produto)
			return
		}
	}
	c.String(404, "Produto não encontrado.")
}

func HandleRequest() {
	r := gin.Default()
	r.GET("/produtos", GetAll)
	r.GET("/produtos/:id", SearchProduct)
	r.Run(":8000")
}

func main() {
	HandleRequest()
}
