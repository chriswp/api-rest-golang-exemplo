package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Venda struct {
	ID              string    `json:"id" gorm:"type:uuid;primary_key"`
	DataVenda       time.Time `json:"data_venda" valid:"notnull"`
	Cliente         Cliente   `valid:"notnull"`
	ClienteID       string
	VendaItem       []*VendaItem `json:"venda_item"`
	Total           float64
	Subtotal        float64
	Observacao      string
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func NewVenda(cliente Cliente, itens []*VendaItem) (*Venda, error) {
	venda := Venda{
		Cliente:   cliente,
		VendaItem: itens,
	}
	venda.prepare()
	err := venda.Validate()

	if err != nil {
		return nil, err
	}

	return &venda, nil
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
