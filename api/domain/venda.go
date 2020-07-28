package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Venda struct {
	ID              string      `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	DataVenda       time.Time   `json:"data_venda" valid:"-"`
	Cliente         *Cliente    `json:"cliente" valid:"required"`
	ClienteID       string      `valid:"-"`
	Itens           []VendaItem `json:"venda_item" valid:"-"`
	Total           float64     `valid:"-"`
	Subtotal        float64     `valid:"-"`
	Observacao      string      `valid:"type(string),optional"`
	DataCriacao     time.Time   `json:"data_criacao" valid:"-"`
	DataModificacao time.Time   `json:"data_modificacao" valid:"-"`
}

func NewVenda(cliente *Cliente, dataVenda time.Time, observacao string) (*Venda, error) {
	venda := Venda{
		Cliente:    cliente,
		DataVenda:  dataVenda,
		Observacao: observacao,
	}
	venda.prepare()
	err := venda.Validate()

	if err != nil {
		return nil, err
	}

	return &venda, nil
}

func (venda *Venda) AddNewItem(produto *Produto, quantidade int) {
	item := VendaItem{Produto: produto, Quantidade: quantidade}
	venda.Itens = append(venda.Itens, item)
}

func (v *Venda) prepare() {
	v.ID = uuid.NewV4().String()
	v.DataVenda = time.Now()
	v.DataCriacao = time.Now()
	v.DataModificacao = time.Now()
}

func (v Venda) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	return err
}
