package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CategoriaRepository interface {
	Insert(categoria *domain.Categoria) (*domain.Categoria, error)
	Find(id string) (*domain.Categoria, error)
}

type CategoriaRepostioryDb struct {
	Db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepostioryDb {
	return &CategoriaRepostioryDb{Db: db}
}

func (repo CategoriaRepostioryDb) Insert(categoria *domain.Categoria) (*domain.Categoria, error) {
	err := repo.Db.Create(categoria).Error
	if err != nil {
		return nil, err
	}

	return categoria, nil
}

func (repo CategoriaRepostioryDb) Find(id string) (*domain.Categoria, error) {
	var categoria domain.Categoria
	repo.Db.Preload("Produtos").First(&categoria, "id = ?", id)

	if categoria.ID == "" {
		return nil, fmt.Errorf("categoria inexistente")
	}

	return &categoria, nil
}
