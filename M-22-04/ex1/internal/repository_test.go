package products_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/junior/products-api/internal/products"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	produtos, ok := data.(*[]products.Product)
	if !ok {
		return fmt.Errorf("Erro: esperava-se um *[]Produto")
	}

	*produtos = []products.Product{
		{ID: 1, Nome: "Jardineira Amarela", Cor: "Amarela", Preco: 79.99, Estoque: 20, Codigo: "JAR789", Tamanho: "P", Publicacao: true, DataCriacao: time.Date(2024, time.April, 11, 18, 07, 26, 532546000, time.FixedZone("GMT", -3*60*60))},
		{ID: 2, Nome: "Jardineira Amarela", Cor: "Amarela", Preco: 79.99, Estoque: 20, Codigo: "JAR789", Tamanho: "P", Publicacao: true, DataCriacao: time.Date(2024, time.April, 11, 18, 12, 01, 332271000, time.FixedZone("GMT", -3*60*60))},
	}
	return nil
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	expectedProducts := []products.Product{
		{ID: 1, Nome: "Jardineira Amarela", Cor: "Amarela", Preco: 79.99, Estoque: 20, Codigo: "JAR789", Tamanho: "P", Publicacao: true, DataCriacao: time.Date(2024, time.April, 11, 18, 07, 26, 532546000, time.FixedZone("GMT", -3*60*60))},
		{ID: 2, Nome: "Jardineira Amarela", Cor: "Amarela", Preco: 79.99, Estoque: 20, Codigo: "JAR789", Tamanho: "P", Publicacao: true, DataCriacao: time.Date(2024, time.April, 11, 18, 12, 01, 332271000, time.FixedZone("GMT", -3*60*60))},
	}

	store := &StubStore{}
	repo := products.NewRepository(store)

	products, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
}
