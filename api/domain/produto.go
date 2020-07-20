package domain

import "time"

type Produto struct {
	ID          string
	Nome        string
	Descricao   string
	Preco       float64
	CategoriaID *Categoria
	Ativo       bool
	DataCriacao time.Time
}
