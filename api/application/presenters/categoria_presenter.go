package presenters

import "api/domain"

type CategoriaPresenter interface {
	ResponseAll(categorias []*domain.Categoria) []*domain.Categoria
	Response(categoria *domain.Categoria) *domain.Categoria
}

type categoriaPresenter struct {
}

func NewCategoriaPresenter() CategoriaPresenter {
	return &categoriaPresenter{}
}

func (cp categoriaPresenter) ResponseAll(categorias []*domain.Categoria) []*domain.Categoria {
	return categorias
}

func (cp categoriaPresenter) Response(categoria *domain.Categoria) *domain.Categoria {
	return categoria
}
