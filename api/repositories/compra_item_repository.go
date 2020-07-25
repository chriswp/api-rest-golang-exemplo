package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CompraItemRepository interface {
	Insert(compra *domain.CompraItem) (*domain.CompraItem, error)
	Find(id string) (*domain.CompraItem, error)
}

type CompraItemRepositoryDb struct {
	Db *gorm.DB
}

func NewCompraItemRepository(db *gorm.DB) *CompraItemRepositoryDb {
	return &CompraItemRepositoryDb{Db: db}
}

func (repo CompraItemRepositoryDb) Insert(item *domain.CompraItem) (*domain.CompraItem, error) {
	err := repo.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (repo CompraItemRepositoryDb) Find(id string) (*domain.CompraItem, error) {
	var compraItem domain.CompraItem
	repo.Db.Preload("Compra").Preload("Produto").First(&compraItem, "id = ?", id)

	if compraItem.ID == "" {
		return nil, fmt.Errorf("item  da compra inexistente")
	}

	return &compraItem, nil
}
