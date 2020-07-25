package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Produto struct {
	ID            string           `json:"id" gorm:"type:uuid;primary_key"`
	Nome          string           `json:"nome" gorm:"type:varchar(20)"`
	Descricao     string           `json:"descricao" valid:"notnull" gorm:"type:varchar(255)"`
	Preco         float64          `json:"preco" valid:"notnull"`
	Categoria     Categoria        `json:"categoria" valid:"-"`
	CategoriaID   string           `json:"-" valid:"-" gorm:"column:categoria_id;type:uuid;notnull"`
	Ativo         bool             `json:"ativo"`
	DataCriacao   time.Time        `json:"data_criacao" valid:"-"`
	ProdutoDigial []*ProdutoDigial `json:"produto_digital" valid:"-" gorm:"ForeignKey:ProdutoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduto(categoria Categoria) (*Produto, error) {
	produto := Produto{
		Categoria: categoria,
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
}
