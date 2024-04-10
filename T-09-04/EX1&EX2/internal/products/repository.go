package products

type Repository interface {
	GetAll() ([]Product, error)
	Store(nome, cor, codigo, tamanho string, estoque int, preco float64, publicacao bool) (Product, error)
	GetByID(id int) (Product, error)
	Create(produto Product) error
	Update(id int, produto Product) (Product, error)
	Delete(id int) error
}

func NewRepository() Repository {
	return &MemoryRepository{}
}
