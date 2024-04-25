package products

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/junior/products-api/internal/products"
	"github.com/junior/products-api/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	input := []products.Product{
		{
			ID:          1,
			Nome:        "Before Update",
			Cor:         "Amarela",
			Preco:       20.20,
			Estoque:     15,
			Codigo:      "JAR123",
			Tamanho:     "G",
			Publicacao:  false,
			DataCriacao: "2006-01-02 15:04:05",
		},
	}

	dataJson, _ := json.Marshal(input)

	dbMock := store.FileMock{
		Data: dataJson,
	}

	myRepo := products.NewRepository(&dbMock)

	// excluindo um produto existente.
	err := myRepo.Delete(1)
	assert.NoError(t, err)

	// excluindo um produto inexistente.
	err = myRepo.Delete(13)
	assert.Error(t, err)

	resp, _ := myRepo.GetAll()

	expected := []products.Product{
		{
			ID:          2,
			Nome:        "Laptop",
			Cor:         "Preto",
			Preco:       1500.99,
			Estoque:     8,
			Codigo:      "LT009",
			Tamanho:     "15.6\"",
			Publicacao:  true,
			DataCriacao: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	// verificando se o produto deletado ainda existe.
	assert.NotEqual(t, expected, resp)
}
