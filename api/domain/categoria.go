package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Categoria struct {
	ID              string     `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Nome            string     `json:"nome" valid:"notnull" gorm:"type:varchar(40);notnull"`
	Ativo           bool       `json:"ativo" valid:"-"`
	DataCriacao     time.Time  `json:"data_criacao" valid:"-"`
	DataModificacao time.Time  `json:"data_modificacao" valid:"-"`
	Produtos        []*Produto `json:"produtos" valid:"-" gorm:"ForeignKey:CategoriaID"`
}

func (Categoria) TableName() string {
	return "categorias"
}

func (categoria *Categoria) BeforeCreate() (err error) {
	categoria.prepare()
	return
}

func (categoria *Categoria) BeforeSave() (err error) {
	categoria.DataModificacao = time.Now()
	return
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
	categoria.DataModificacao = time.Now()
	categoria.Ativo = true
}

func (categoria Categoria) Validate() error {
	_, err := govalidator.ValidateStruct(categoria)
	return err
}
