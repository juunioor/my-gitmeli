package products

type Product struct {
	ID          int     `json:"id"`
	Nome        string  `json:"nome"`
	Cor         string  `json:"cor"`
	Preco       float64 `json:"preco"`
	Estoque     int     `json:"estoque"`
	Codigo      string  `json:"codigo"`
	Tamanho     string  `json:"tamanho"`
	Publicacao  bool    `json:"publicacao"`
	DataCriacao string  `json:"data_criacao"`
}
