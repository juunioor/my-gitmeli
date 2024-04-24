package products

import (
	"fmt"
	"time"

	"github.com/junior/products-api/pkg/store"
)

type FileRepository struct {
	db store.Store
}

func NewFileRepository(db store.Store) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) GetByID(id int) (Product, error) {
	var p Product
	return p, nil
}

func (r *FileRepository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	fmt.Println("Products retrieved:", ps)
	return ps, nil
}

func (r *FileRepository) Delete(id int) error {
	var ps []Product
	r.db.Read(&ps)
	index := -1
	for i, p := range ps {
		if p.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("produto não encontrado")
	}

	ps = append(ps[:index], ps[index+1:]...)
	fmt.Println("Products after deletion:", ps)

	r.db.Write(&ps)

	return nil
}

func (r *FileRepository) Update(id int, nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
	var ps []Product
	r.db.Read(&ps)

	p := Product{
		Nome:        nome,
		Cor:         cor,
		Preco:       preco,
		Estoque:     estoque,
		Codigo:      codigo,
		Tamanho:     tamanho,
		Publicacao:  publicacao,
		DataCriacao: time.Now().Format("2006-01-02 15:04:05"),
	}

	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p

			r.db.Write(&ps)
			return p, nil
		}
	}

	return Product{}, fmt.Errorf("produto não encontrado")
}

func (r *FileRepository) UpdateFields(id int, nome string, preco float64) (Product, error) {
	var ps []Product
	var p Product
	r.db.Read(&ps)

	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			if nome != "" {
				ps[i].Nome = nome
			}

			if preco != 0 {
				ps[i].Preco = preco
			}
			p = ps[i]

			return p, nil
		}
	}
	return Product{}, nil
}

func (r *FileRepository) Create(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
	p := Product{
		ID:          0,
		Nome:        nome,
		Cor:         cor,
		Preco:       preco,
		Estoque:     estoque,
		Codigo:      codigo,
		Tamanho:     tamanho,
		Publicacao:  publicacao,
		DataCriacao: time.Now().Format("2006-01-02 15:04:05"),
	}

	var ps []Product
	r.db.Read(&ps)

	lastID := len(ps) + 1
	p.ID = int(lastID)
	ps = append(ps, p)
	err := r.db.Write(ps)
	if err != nil {
		return Product{}, err
	}

	return p, err
}
