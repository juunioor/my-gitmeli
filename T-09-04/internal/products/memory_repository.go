package products

import "time"

var ps []Product
var p Product
var lastID int = 0

type MemoryRepository struct {
}

func (m *MemoryRepository) GetAll() ([]Product, error) {
	return ps, nil
}

func (m *MemoryRepository) GetByID(id int) (Product, error) {
	return p, nil
}

func (m *MemoryRepository) Create(produto Product) error {
	return nil
}

func (m *MemoryRepository) Update(id int, produto Product) (Product, error) {
	return produto, nil
}

func (m *MemoryRepository) Delete(id int) error {
	return nil
}

func (m *MemoryRepository) LastID() (int, error) {
	return lastID, nil
}

func (m *MemoryRepository) Store(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
	lastID++
	p := Product{
		ID:          lastID,
		Nome:        nome,
		Preco:       preco,
		Estoque:     estoque,
		Codigo:      codigo,
		Tamanho:     tamanho,
		Publicacao:  publicacao,
		DataCriacao: time.Now(),
	}
	ps = append(ps, p)
	return p, nil
}
