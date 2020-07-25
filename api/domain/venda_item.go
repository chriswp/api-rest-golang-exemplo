package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type VendaItem struct {
	ID              string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Produto         *Produto  `json:"produto" valid:"required"`
	ProdutoID       string    `json:"-" valid:"-"`
	Venda           *Venda    `json:"venda" valid:"required"`
	VendaID         string    `json:"-"  valid:"-"`
	Quantidade      int       `valid:"notnull"`
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func (VendaItem) TableName() string {
	return "venda_itens"
}

func NewVendaItem(venda *Venda, produto *Produto, quantidade int) (*VendaItem, error) {
	vendaItem := VendaItem{
		Produto:    produto,
		Venda:      venda,
		Quantidade: quantidade,
	}

	vendaItem.prepare()
	err := vendaItem.Validate()

	if err != nil {
		return nil, err
	}

	return &vendaItem, nil
}

func (vt *VendaItem) prepare() {
	vt.ID = uuid.NewV4().String()
	vt.DataCriacao = time.Now()
	vt.DataModificacao = time.Now()
}

func (vi VendaItem) Validate() error {
	_, err := govalidator.ValidateStruct(vi)
	return err
}
