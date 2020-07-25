package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewProduto(t *testing.T) {
	categoria, _ := domain.NewCategoria("categoria")

	produto, err := domain.NewProduto(categoria, "produto", 10, "apenas para teste")
	require.NotNil(t, produto)
	require.Nil(t, err)
}

func TestNewProdutoFieldsEmpty(t *testing.T) {
	categoria, _ := domain.NewCategoria("categoria")

	_, err := domain.NewProduto(categoria, "", 10.52, "")
	require.Error(t, err)

	_, nomeEmpty := domain.NewProduto(categoria, "", 10.52, "apenas para teste")
	require.Error(t, nomeEmpty)

	_, descricaoEmpty := domain.NewProduto(categoria, "produto", 10, "")
	require.Error(t, descricaoEmpty)
}
