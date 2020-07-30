package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProdutoRepository interface {
	Insert(produto *domain.Produto) (*domain.Produto, error)
	Find(id string) (*domain.Produto, error)
	All() ([]*domain.Produto, error)
}

type produtoRepostiory struct {
	Db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) ProdutoRepository {
	return &produtoRepostiory{Db: db}
}

func (repo produtoRepostiory) Insert(produto *domain.Produto) (*domain.Produto, error) {
	err := repo.Db.Create(produto).Error
	if err != nil {
		return nil, err
	}

	return produto, nil
}

func (repo produtoRepostiory) Find(id string) (*domain.Produto, error) {
	var produto domain.Produto
	repo.Db.Preload("ProdutoDigital").First(&produto, "id = ?", id)

	if produto.ID == "" {
		return nil, fmt.Errorf("produto inexistente")
	}

	return &produto, nil
}

func (repo produtoRepostiory) All() ([]*domain.Produto, error) {
	var produtos []*domain.Produto
	err := repo.Db.Find(&produtos).Error

	if err != nil {
		return nil, err
	}

	return produtos, nil
}
