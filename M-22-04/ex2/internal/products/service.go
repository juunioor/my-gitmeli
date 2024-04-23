package products

import (
	"errors"
)

type Service interface {
	GetAll() ([]Product, error)
	Create(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error)
	GetByID(id int) (Product, error)
	Update(id int, nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error)
	UpdateFields(id int, nome string, preco float64) (Product, error)
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

func (s *service) Create(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {

	product, err := s.repository.Create(nome, cor, codigo, tamanho, estoque, preco, publicacao)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error) {
	return s.repository.Update(id, nome, cor, codigo, tamanho, estoque, preco, publicacao)
}

func (s *service) UpdateFields(id int, nome string, preco float64) (Product, error) {
	return s.repository.UpdateFields(id, nome, preco)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
