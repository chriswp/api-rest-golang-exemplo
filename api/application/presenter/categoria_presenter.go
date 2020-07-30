package presenter

import "api/domain"

type CategoriaPresenter interface {
	Response(categorias []*domain.Categoria) []*domain.Categoria
}

type categoriaPresenter struct {
}

func NewCategoriaPresenter() CategoriaPresenter {
	return &categoriaPresenter{}
}

func (cp categoriaPresenter) Response(categorias []*domain.Categoria) []*domain.Categoria {
	for _, c := range categorias {
		c.ID = c.ID
		c.Nome = c.Nome
	}
	return categorias
}
