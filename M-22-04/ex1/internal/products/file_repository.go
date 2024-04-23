package products

import (
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
	return p, nil
}

func (r *FileRepository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *FileRepository) Delete(id int) error {
	return nil
}

func (r *FileRepository) Update(id int, nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
	return Product{}, nil
}
func (r *FileRepository) UpdateFields(id int, nome string, preco float64) (Product, error) {
	return Product{}, nil
}

func (r *FileRepository) Create(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
	p := Product{
		ID:          lastID,
		Nome:        nome,
		Cor:         cor,
		Preco:       preco,
		Estoque:     estoque,
		Codigo:      codigo,
		Tamanho:     tamanho,
		Publicacao:  publicacao,
		DataCriacao: time.Now(),
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
