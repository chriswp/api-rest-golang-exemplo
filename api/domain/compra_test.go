package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewCompra(t *testing.T) {
	compra, err := domain.NewCompra(time.Now())
	require.NotNil(t, compra)
	require.Nil(t, err)
}

func TestCompra_AddNewItem(t *testing.T) {
	compra, err := domain.NewCompra(time.Now())

	categoria, _ := domain.NewCategoria("categoria teste")
	produto, _ := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")

	compra.AddNewItem(produto, 2)
	require.NotNil(t, compra)
	require.Equal(t, len(compra.Itens), 1)
	require.Nil(t, err)

	compra.AddNewItem(produto, 1)
	require.Equal(t, len(compra.Itens), 2)
	require.Nil(t, err)
}
