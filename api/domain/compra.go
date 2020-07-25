package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Compra struct {
	ID              string    `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	DataCompra      time.Time `valid:"notnull"`
	Observacao      string
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
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

func (c *Compra) prepare() {
	c.ID = uuid.NewV4().String()
	c.DataCriacao = time.Now()
	c.DataModificacao = time.Now()
}

func (c *Compra) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
