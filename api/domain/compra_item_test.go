package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewCompraItem(t *testing.T) {
	compra, _ := domain.NewCompra(time.Now())
	categoria, _ := domain.NewCategoria("categoria teste")

	produtos := []*domain.Produto{}
	produto, err := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")
	produtos = append(produtos, produto)

	item, err := domain.NewCompraItem(compra, produtos, 10)
	require.NotNil(t, item)
	require.Nil(t, err)
}
