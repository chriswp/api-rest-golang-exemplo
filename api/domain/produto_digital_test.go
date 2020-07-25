package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewProdutoDigital(t *testing.T) {
	categoria, _ := domain.NewCategoria("categoria teste")
	produto, _ := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")

	pDigital, err := domain.NewProdutoDigital(produto, "http://link.exemplo.com")
	require.NotNil(t, pDigital)
	require.Nil(t, err)
}
