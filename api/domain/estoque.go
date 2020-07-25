package domain

import "time"

type Estoque struct {
	ID                   string     `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	Produto              []*Produto `json:"-" valid:"notnull"`
	ProdutoID            string     `json:"produto_id"`
	Data                 time.Time
	QuantidadeDisponivel int8
	DataCriacao          time.Time `json:"data_criacao" valid:"-"`
	DataModificacao      time.Time `json:"data_modificacao" valid:"-"`
}
