package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProdutoDigitalRepository interface {
	Insert(pDigital *domain.ProdutoDigital) (*domain.ProdutoDigital, error)
	Find(id string) (*domain.ProdutoDigital, error)
}

type ProdutoDigitalRepostioryDb struct {
	Db *gorm.DB
}

func NewProdutoDigitalRepository(db *gorm.DB) *ProdutoDigitalRepostioryDb {
	return &ProdutoDigitalRepostioryDb{Db: db}
}

func (repo ProdutoDigitalRepostioryDb) Insert(pDigital *domain.ProdutoDigital) (*domain.ProdutoDigital, error) {
	err := repo.Db.Create(pDigital).Error
	if err != nil {
		return nil, err
	}

	return pDigital, nil
}

func (repo ProdutoDigitalRepostioryDb) Find(id string) (*domain.ProdutoDigital, error) {
	var pDigital domain.ProdutoDigital
	repo.Db.Preload("Produto").First(&pDigital, "id = ?", id)

	if pDigital.ID == "" {
		return nil, fmt.Errorf("Produto Digital inexistente")
	}

	return &pDigital, nil
}
