package repositories

import "api/domain"

type CategoriaRepository interface {
	Find(id string) (*domain.Categoria, error)
	Insert(categoria *domain.Categoria) (*domain.Categoria, error)
	All(categorias []*domain.Categoria) ([]*domain.Categoria, error)
	Update(categoria *domain.Categoria, id string) (*domain.Categoria, error)
}
