package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cliente struct {
	ID              string `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	CPF             string
	email           string
	Vendas          []*Venda  `json:"vendas" gorm:"ForeignKey:ClienteID"`
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func NewCliente() (*Cliente, error) {
	cliente := Cliente{}
	cliente.prepare()
	err := cliente.Validate()

	if err != nil {
		return nil, err
	}

	return &cliente, nil
}

func (c *Cliente) prepare() {
	c.ID = uuid.NewV4().String()
	c.DataCriacao = time.Now()
	c.DataModificacao = time.Now()
}

func (c Cliente) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
