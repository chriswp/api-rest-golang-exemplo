package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Produto struct {
	ID              string          `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Nome            string          `json:"nome" valid:"notnull" gorm:"type:varchar(20);notnull"`
	Descricao       string          `json:"descricao" valid:"notnull" gorm:"type:varchar(255)"`
	Preco           float64         `json:"preco" valid:"type(float64),notnull" gorm:"notnull"`
	Categoria       *Categoria      `json:"categoria" valid:"required" gorm:"foreignkey:CategoriaID"`
	CategoriaID     string          `json:"-" valid:"-" gorm:"column:categoria_id;type:uuid;notnull"`
	Ativo           bool            `json:"ativo" valid:"required"`
	DataCriacao     time.Time       `json:"data_criacao" valid:"-"`
	DataModificacao time.Time       `json:"data_modificacao" valid:"-"`
	ProdutoDigital  *ProdutoDigital `json:"produto_digital" valid:"-" gorm:"ForeignKey:ProdutoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduto(categoria *Categoria, nome string, preco float64, descricao string) (*Produto, error) {
	produto := Produto{
		Nome:      nome,
		Categoria: categoria,
		Descricao: descricao,
		Preco:     preco,
	}
	produto.prepare()
	err := produto.Validate()

	if err != nil {
		return nil, err
	}

	return &produto, nil
}

func (produto Produto) Validate() error {
	_, err := govalidator.ValidateStruct(produto)
	return err
}

func (p *Produto) prepare() {
	p.ID = uuid.NewV4().String()
	p.DataCriacao = time.Now()
	p.DataModificacao = time.Now()
	p.Ativo = true
}
