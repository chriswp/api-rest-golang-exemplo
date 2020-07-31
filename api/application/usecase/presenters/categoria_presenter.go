package presenters

import "api/domain"

type CategoriaPresenter interface {
	Response(categorias []*domain.Categoria) []*domain.Categoria
}
