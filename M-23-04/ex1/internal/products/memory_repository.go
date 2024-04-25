package products

// import (
// 	"fmt"
// 	"time"
// )

// var ps []Product
// var p Product
// var lastID int = 0

// type MemoryRepository struct {
// }

// func (m *MemoryRepository) GetAll() ([]Product, error) {
// 	return ps, nil
// }

// func (m *MemoryRepository) GetByID(id int) (Product, error) {
// 	return p, nil
// }

// func (m *MemoryRepository) Update(id int, nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
// 	p := Product{
// 		Nome:        nome,
// 		Cor:         cor,
// 		Preco:       preco,
// 		Estoque:     estoque,
// 		Codigo:      codigo,
// 		Tamanho:     tamanho,
// 		Publicacao:  publicacao,
// 		DataCriacao: time.Now(),
// 	}

// 	for i := range ps {
// 		if ps[i].ID == id {
// 			p.ID = id
// 			ps[i] = p

// 			return p, nil
// 		}
// 	}

// 	return Product{}, fmt.Errorf("produto não encontrado")
// }

// func (m *MemoryRepository) UpdateFields(id int, nome string, preco float64) (Product, error) {

// 	for i := range ps {
// 		if ps[i].ID == id {
// 			p.ID = id
// 			if nome != "" {
// 				ps[i].Nome = nome
// 			}

// 			if preco != 0 {
// 				ps[i].Preco = preco
// 			}
// 			p = ps[i]

// 			return p, nil
// 		}
// 	}

// 	return Product{}, fmt.Errorf("produto não encontrado")
// }

// func (m *MemoryRepository) Delete(id int) error {

// 	index := -1
// 	for i := range ps {
// 		if ps[i].ID == id {
// 			index = i
// 		}
// 	}
// 	if index == -1 {
// 		return fmt.Errorf("produto não encontrado")
// 	}

// 	ps = append(ps[:index], ps[index+1:]...)
// 	return nil
// }

// func (m *MemoryRepository) LastID() (int, error) {
// 	return lastID, nil
// }

// func (m *MemoryRepository) Create(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
// 	lastID++
// 	p := Product{
// 		ID:          lastID,
// 		Nome:        nome,
// 		Cor:         cor,
// 		Preco:       preco,
// 		Estoque:     estoque,
// 		Codigo:      codigo,
// 		Tamanho:     tamanho,
// 		Publicacao:  publicacao,
// 		DataCriacao: time.Now(),
// 	}
// 	ps = append(ps, p)
// 	return p, nil
// }
