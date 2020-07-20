package domain

import "time"

type Categoria struct {
	ID          string
	Nome        string
	DataCriacao time.Time
	Ativo       bool
}
