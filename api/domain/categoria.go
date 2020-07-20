package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Categoria struct {
	ID          string    `json:"id"  gorm:"type:uuid;primary_key"`
	Nome        string    `json:"nome" valid:"notnull"`
	DataCriacao time.Time `json:"data_criacao" valid:"-"`
	Ativo       bool      `json:"ativo" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCategoria() (*Categoria, error) {
	categoria := Categoria{}
	categoria.prepare()

	err := categoria.Validate()

	if err != nil {
		return nil, err
	}

	return &categoria, nil
}

func (categoria *Categoria) prepare() {
	categoria.ID = uuid.NewV4().String()
	categoria.DataCriacao = time.Now()
	categoria.Ativo = true
}

func (categoria Categoria) Validate() error {
	_, err := govalidator.ValidateStruct(categoria)
	return err
}
