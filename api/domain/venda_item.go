package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VendaItem struct {
	ID              string    `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	Produto         Produto   `json:"produto" valid:"notnull"`
	ProdutoID       string    `json:"-"`
	Venda           Venda     `json:"venda" valid:"notnull"`
	VendaID         Venda     `json:"-"`
	Quantidade      int8      `valid:"notnull"`
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func NewVendaItem(venda Venda, produto Produto, quantidade int8) (*VendaItem, error) {
	vendaItem := VendaItem{
		Produto:    produto,
		Venda:      venda,
		Quantidade: quantidade,
	}

	return &vendaItem, nil
}

func (vt *VendaItem) prepare() {
	vt.ID = uuid.NewV4().String()
	vt.DataCriacao = time.Now()
	vt.DataModificacao = time.Now()
}
