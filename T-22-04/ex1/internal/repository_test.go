package products

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/junior/products-api/internal/products"
	"github.com/junior/products-api/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
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

	// update em um produto existente.
	product, err := myRepo.Update(1, "After Update", "Amarela", "JAR123", "G", 15, 20.20, false)
	assert.NoError(t, err)

	expected := products.Product{
		ID:          1,
		Nome:        "After Update",
		Cor:         "Amarela",
		Preco:       20.20,
		Estoque:     15,
		Codigo:      "JAR123",
		Tamanho:     "G",
		Publicacao:  false,
		DataCriacao: time.Now().Format("2006-01-02 15:04:05"),
	}

	// verificando se o produto deletado ainda existe.
	assert.Equal(t, expected, product)
}
