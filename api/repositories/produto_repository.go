package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProdutoRepository interface {
	Insert(produto *domain.Produto) (*domain.Produto, error)
	Find(id string) (*domain.Produto, error)
}

type ProdutoRepostioryDb struct {
	Db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) *ProdutoRepostioryDb {
	return &ProdutoRepostioryDb{Db: db}
}

func (repo ProdutoRepostioryDb) Insert(produto *domain.Produto) (*domain.Produto, error) {
	err := repo.Db.Create(produto).Error
	if err != nil {
		return nil, err
	}

	return produto, nil
}

func (repo ProdutoRepostioryDb) Find(id string) (*domain.Produto, error) {
	var produto domain.Produto
	repo.Db.Preload("ProdutoDigital").First(&produto, "id = ?", id)

	if produto.ID == "" {
		return nil, fmt.Errorf("produto inexistente")
	}

	return &produto, nil
}
