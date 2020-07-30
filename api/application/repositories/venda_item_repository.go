package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type VendaItemRepository interface {
	Insert(item *domain.VendaItem) (*domain.VendaItem, error)
	Find(id string) (*domain.VendaItem, error)
}

type VendaItemRepostioryDb struct {
	Db *gorm.DB
}

func NewVendaItemRepository(db *gorm.DB) *VendaItemRepostioryDb {
	return &VendaItemRepostioryDb{Db: db}
}

func (repo VendaItemRepostioryDb) Insert(item *domain.VendaItem) (*domain.VendaItem, error) {
	err := repo.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (repo VendaItemRepostioryDb) Find(id string) (*domain.VendaItem, error) {
	var item domain.VendaItem
	repo.Db.Preload("Produto").Preload("Venda").First(&item, "id = ?", id)

	if item.ID == "" {
		return nil, fmt.Errorf("item inexistente")
	}

	return &item, nil
}
