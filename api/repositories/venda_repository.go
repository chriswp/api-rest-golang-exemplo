package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type VendaRepository interface {
	Insert(venda *domain.Venda) (*domain.Venda, error)
	Find(id string) (*domain.Venda, error)
}

type VendaRepostioryDb struct {
	Db *gorm.DB
}

func NewVendaRepository(db *gorm.DB) *VendaRepostioryDb {
	return &VendaRepostioryDb{Db: db}
}

func (repo VendaRepostioryDb) Insert(venda *domain.Venda) (*domain.Venda, error) {
	err := repo.Db.Create(venda).Error
	if err != nil {
		return nil, err
	}

	return venda, nil
}

func (repo VendaRepostioryDb) Find(id string) (*domain.Venda, error) {
	var venda domain.Venda
	repo.Db.Preload("Itens").First(&venda, "id = ?", id)

	if venda.ID == "" {
		return nil, fmt.Errorf("venda inexistente")
	}

	return &venda, nil
}
