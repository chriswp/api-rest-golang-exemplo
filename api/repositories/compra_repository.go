package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CompraRepository interface {
	Insert(compra *domain.Compra) (*domain.Compra, error)
	Find(id string) (*domain.Compra, error)
}

type CompraRepositoryDb struct {
	Db *gorm.DB
}

func NewCompraRepository(db *gorm.DB) *CompraRepositoryDb {
	return &CompraRepositoryDb{Db: db}
}

func (repo CompraRepositoryDb) Insert(compra *domain.Compra) (*domain.Compra, error) {
	err := repo.Db.Create(compra).Error
	if err != nil {
		return nil, err
	}

	return compra, nil
}

func (repo CompraRepositoryDb) Find(id string) (*domain.Compra, error) {
	var compra domain.Compra
	repo.Db.First(&compra, "id = ?", id)

	if compra.ID == "" {
		return nil, fmt.Errorf("compra inexistente")
	}

	return &compra, nil
}
