package products_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/junior/products-api/internal/products"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	ReadCalled bool
}

func (s *StubStore) Read(data interface{}) error {
	s.ReadCalled = true
	produtos, ok := data.(*[]products.Product)
	if !ok {
		return fmt.Errorf("Erro: esperava-se um *[]Produto")
	}

	*produtos = []products.Product{
		{ID: 1, Nome: "Before Update", Cor: "Amarela", Preco: 79.99, Estoque: 20, Codigo: "JAR789", Tamanho: "P", Publicacao: true, DataCriacao: time.Date(2024, time.April, 11, 18, 07, 26, 532546000, time.FixedZone("GMT", -3*60*60))},
	}
	return nil
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	expectedProducts := products.Product{
		ID: 1, Nome: "After Update", Cor: "Amarela", Preco: 70.20, Estoque: 20, Codigo: "JAR789", Tamanho: "P", Publicacao: true, DataCriacao: time.Date(2024, time.April, 11, 18, 07, 26, 532546000, time.FixedZone("GMT", -3*60*60)),
	}

	store := &StubStore{}
	repo := products.NewRepository(store)

	products, err := repo.UpdateFields(1, "After Update", 70.20)
	assert.NoError(t, err)

	assert.Equalf(t, expectedProducts, products, "O produto recebido não é igual.")
	assert.Truef(t, store.ReadCalled, "O método Read não foi chamado.")
}
