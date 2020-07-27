package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Compra struct {
	ID              string       `json:"id" valid:"uuid" gorm:"type:uuid;primary_key;not null"`
	DataCompra      time.Time    `valid:"required" gorm:"type:date"`
	Observacao      string       `valid:"type(string),optional "gorm:"type:text"`
	DataCriacao     time.Time    `json:"data_criacao" valid:"-"`
	DataModificacao time.Time    `json:"data_modificacao" valid:"-"`
	Itens           []CompraItem `json:"itens" valid:"-" gorm:"foreignkey:CompraID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCompra(dataCompra time.Time) (*Compra, error) {
	compra := Compra{
		DataCompra: dataCompra,
	}
	compra.prepare()

	err := compra.Validate()
	if err != nil {
		return nil, err
	}

	return &compra, nil
}

func (compra *Compra) AddNewItem(produto *Produto, quantidade int) {
	item := CompraItem{Produto: produto, Quantidade: quantidade}
	compra.Itens = append(compra.Itens, item)
}

func (c *Compra) prepare() {
	c.ID = uuid.NewV4().String()
	c.DataCriacao = time.Now()
	c.DataModificacao = time.Now()
}

func (c *Compra) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
