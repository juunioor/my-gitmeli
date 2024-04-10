package products

import (
	"errors"
)

type Service interface {
	GetAll() ([]Product, error)
	Store(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error)
	GetByID(id int) (Product, error)
	Create(produto Product) error
	Update(id int, produto Product) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	produtos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func (s *service) GetByID(id int) (Product, error) {
	produtos, err := s.repository.GetAll()
	if err != nil {
		return Product{}, errors.New("falha ao obter produtos")
	}

	for _, produto := range produtos {
		if produto.ID == id {
			return produto, nil
		}
	}

	return Product{}, errors.New("produto não encontrado")
}

func (s *service) Store(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {

	product, err := s.repository.Store(nome, cor, codigo, tamanho, estoque, preco, publicacao)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Create(produto Product) error {
	_, err := s.repository.GetByID(produto.ID)
	if err != nil {
		return errors.New("produto já existente")
	}

	_, err = s.repository.Store(produto.Nome, produto.Cor, produto.Codigo, produto.Tamanho, produto.Estoque, produto.Preco, produto.Publicacao)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(id int, produto Product) (Product, error) {
	_, err := s.repository.GetByID(id)
	if err != nil {
		return Product{}, errors.New("produto não encontrado")
	}

	_, err = s.repository.Update(id, produto)
	if err != nil {
		return Product{}, err
	}
	return Product{}, nil
}

func (s *service) Delete(id int) error {
	_, err := s.repository.GetByID(id)
	if err != nil {
		return errors.New("produto não encontrado")
	}

	err = s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
