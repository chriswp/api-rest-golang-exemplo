package presenters

import "api/domain"

type CategoriaPresenter interface {
	ResponseAll(categorias []*domain.Categoria) []*domain.Categoria
	Response(categoria *domain.Categoria) *domain.Categoria
}
