package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CategoriaRepository interface {
	Find(id string) (*domain.Categoria, error)
	Insert(categoria *domain.Categoria) (*domain.Categoria, error)
	All(categorias []*domain.Categoria) ([]*domain.Categoria, error)
	Update(categoria *domain.Categoria, id string) (*domain.Categoria, error)
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

func (repo categoriaRepository) All(categorias []*domain.Categoria) ([]*domain.Categoria, error) {
	err := repo.Db.Find(&categorias).Error

	if err != nil {
		return nil, err
	}

	return categorias, nil
}

func (repo categoriaRepository) Update(categoria *domain.Categoria, id string) (*domain.Categoria, error) {
	dadosCategoria := domain.Categoria{}
	err := repo.Db.First(&dadosCategoria, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = repo.Db.Model(&dadosCategoria).Updates(map[string]interface{}{
		"Nome":  categoria.Nome,
		"Ativo": categoria.Ativo,
	}).Error
	if err != nil {
		return nil, err
	}

	return &dadosCategoria, nil
}
