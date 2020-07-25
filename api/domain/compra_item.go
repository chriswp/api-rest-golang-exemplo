package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CompraItem struct {
	ID              string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Compra          *Compra   `json:"compra" valid:"required" gorm:"foreignkey:CompraID"`
	CompraID        string    `json:"-" valid:"-" gorm:"column:compra_id;type:uuid;notnull"`
	Produto         *Produto  `json:"produto" valid:"required" gorm:"foreignkey:ProdutoID"`
	ProdutoID       string    `json:"-" valid:"-" gorm:"column:produto_id;type:uuid;notnull"`
	Quantidade      int       `valid:"-"`
	Ativo           bool      `valid:"-"`
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func (CompraItem) TableName() string {
	return "compra_itens"
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCompraItem(compra *Compra, produto *Produto, quantidade int) (*CompraItem, error) {
	itens := CompraItem{
		Compra:     compra,
		Produto:    produto,
		Quantidade: quantidade,
	}
	itens.prepare()
	err := itens.Validate()

	if err != nil {
		return nil, err
	}

	return &itens, nil
}

func (ci *CompraItem) prepare() {
	ci.ID = uuid.NewV4().String()
	ci.DataCriacao = time.Now()
	ci.DataModificacao = time.Now()
	ci.Ativo = true
}

func (ci *CompraItem) Validate() error {
	_, err := govalidator.ValidateStruct(ci)
	return err
}
