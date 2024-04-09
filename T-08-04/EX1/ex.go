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
	Tamanho     string    `json:"tamanho"`
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
		Tamanho:     "M",
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
		Tamanho:     "38",
		Publicacao:  true,
		DataCriacao: time.Now(),
	},
	{
		ID:          3,
		Nome:        "Sapato",
		Cor:         "Vermelho",
		Preco:       67.99,
		Estoque:     56,
		Codigo:      "DLK123",
		Tamanho:     "40",
		Publicacao:  true,
		DataCriacao: time.Now(),
	},
	{
		ID:          4,
		Nome:        "Nike",
		Cor:         "Vermelho",
		Preco:       97.99,
		Estoque:     26,
		Codigo:      "DHG123",
		Tamanho:     "40",
		Publicacao:  true,
		DataCriacao: time.Now(),
	},
}

func SearchProduct(c *gin.Context) {
	cor := c.Query("cor")
	tamanho := c.Query("tamanho")

	var produtosFiltrados []Produto

	for _, produto := range produtos {
		if (produto.Cor != cor && cor != "") || (produto.Tamanho != tamanho && tamanho != "") {
			continue
		}
		produtosFiltrados = append(produtosFiltrados, produto)
	}
	c.JSON(200, produtosFiltrados)
}

func HandleRequest() {
	r := gin.Default()
	r.GET("/produtos", SearchProduct)
	r.Run(":8000")
}

func main() {
	HandleRequest()
}
