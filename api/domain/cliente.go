package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cliente struct {
	ID              string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Nome            string    `json:"nome" valid:"notnull" gorm:"type:varchar(40);notnull"`
	CPF             string    `json:"cpf" valid:"notnull" gorm:"type:varchar(14);unique_index;notnull"`
	Email           string    `json:"email" valid:"email" gorm:"type:varchar(100);unique_index;notnull"`
	Vendas          []*Venda  `json:"vendas" valid:"-" gorm:"ForeignKey:ClienteID"`
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCliente(nome string, cpf string, email string) (*Cliente, error) {
	cliente := Cliente{
		Nome:  nome,
		CPF:   cpf,
		Email: email,
	}
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
