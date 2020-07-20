package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Produto struct {
	ID          string     `json:"id" gorm:"type:uuid;primary_key"`
	Nome        string     `json:"nome" valid:"notnull" gorm:"type:varchar(20)"`
	Descricao   string     `json:"descricao" valid:"notnull" gorm:"type:varchar(255)"`
	Preco       float64    `json:"preco" valid:"notnull"`
	Categoria   *Categoria `json:"categoria" valid:"-"`
	CategoriaID string     `json:"-" valid:"-" gorm:"column:categoria_id;type:uuid;notnull"`
	Ativo       bool       `json:"ativo"`
	DataCriacao time.Time  `json:"data_criacao" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduto() *Produto {
	return &Produto{}
}

func (produto Produto) Validate() error {
	_, err := govalidator.ValidateStruct(produto)
	return err
}