package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CategoriaRepository interface {
	Insert(categoria *domain.Categoria) (*domain.Categoria, error)
	Find(id string) (*domain.Categoria, error)
	All() ([]*domain.Categoria, error)
}

type categoriaRepository struct {
	Db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) CategoriaRepository {
	return &categoriaRepository{Db: db}
}

func (repo categoriaRepository) Insert(categoria *domain.Categoria) (*domain.Categoria, error) {
	err := repo.Db.Create(categoria).Error
	if err != nil {
		return nil, err
	}

	return categoria, nil
}

func (repo categoriaRepository) Find(id string) (*domain.Categoria, error) {
	var categoria domain.Categoria
	repo.Db.Preload("Produtos").First(&categoria, "id = ?", id)

	if categoria.ID == "" {
		return nil, fmt.Errorf("categoria inexistente")
	}

	return &categoria, nil
}

func (repo categoriaRepository) All() ([]*domain.Categoria, error) {
	var categorias []*domain.Categoria
	err := repo.Db.Find(&categorias).Error

	if err != nil {
		return nil, err
	}

	return categorias, nil
}
