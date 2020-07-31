package repositories

import "api/domain"

type CategoriaRepository interface {
	Insert(categoria *domain.Categoria) (*domain.Categoria, error)
	Find(id string) (*domain.Categoria, error)
	All() ([]*domain.Categoria, error)
}
