package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCategoria(t *testing.T) {
	categoria, err := domain.NewCategoria("categoria teste")
	require.NotNil(t, categoria)
	require.Nil(t, err)
}

func TestCategoriaIdIsNotUuid(t *testing.T) {
	categoria := domain.Categoria{}
	categoria.ID = "a1b2c3"
	categoria.Nome = "categoria"
	categoria.Ativo = true

	err := categoria.Validate()
	require.Error(t, err)
}

func TestCategoriaNomeNotNull(t *testing.T) {
	categoria := domain.Categoria{}
	categoria.ID = "a1b2c3"
	categoria.Nome = ""

	err := categoria.Validate()
	require.Error(t, err)
}
