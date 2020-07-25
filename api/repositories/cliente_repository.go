package repositories

import (
	"api/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ClienteRepository interface {
	Insert(cliente *domain.Cliente) (*domain.Cliente, error)
	Find(id string) (*domain.Cliente, error)
}

type ClienteRepositoryDb struct {
	Db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) *ClienteRepositoryDb {
	return &ClienteRepositoryDb{Db: db}
}

func (repo ClienteRepositoryDb) Insert(cliente *domain.Cliente) (*domain.Cliente, error) {
	err := repo.Db.Create(cliente).Error
	if err != nil {
		return nil, err
	}

	return cliente, nil
}

func (repo ClienteRepositoryDb) Find(id string) (*domain.Cliente, error) {
	var cliente domain.Cliente
	repo.Db.Preload("Vendas").First(&cliente, "id = ?", id)

	if cliente.ID == "" {
		return nil, fmt.Errorf("cliente inexistente")
	}

	return &cliente, nil
}
