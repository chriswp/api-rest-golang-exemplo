package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CompraItem struct {
	ID              string     `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	Compra          Compra     `json:"compra" valid:"notnull" gorm:"foreignkey:CompraID"`
	CompraID        string     `json:"-" gorm:"column:compra_id;type:uuid;notnull"`
	Produtos        []*Produto `json:"produto" valid:"notnull" gorm:"foreignkey:ProdutoID"`
	ProdutoID       string     `json:"-" gorm:"column:produto_id;type:uuid;notnull"`
	Quantidade      int        `valid:"notnull"`
	DataCriacao     time.Time  `json:"data_criacao" valid:"-"`
	DataModificacao time.Time  `json:"data_modificacao" valid:"-"`
}

func NewCompraItem(compra Compra, produtos []*Produto, quantidade int) (*CompraItem, error) {
	itens := CompraItem{
		Compra:     compra,
		Produtos:   produtos,
		Quantidade: quantidade,
	}
	itens.prepare()
	err := itens.Validate()

	if err != nil {
		return nil, err
	}

	return &itens, nil
}

func (ci *CompraItem) prepare() {
	ci.ID = uuid.NewV4().String()
	ci.DataCriacao = time.Now()
	ci.DataModificacao = time.Now()
}

func (ci *CompraItem) Validate() error {
	_, err := govalidator.ValidateStruct(ci)
	return err
}
