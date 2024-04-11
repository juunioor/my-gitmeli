package products

type Repository interface {
	GetAll() ([]Product, error)
	Create(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error)
	GetByID(id int) (Product, error)
	Update(id int, nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error)
	UpdateFields(id int, nome string, preco float64) (Product, error)
	Delete(id int) error
}

func NewRepository() Repository {
	return &MemoryRepository{}
}
