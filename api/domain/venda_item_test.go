package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewVendaItem(t *testing.T) {
	cliente, _ := domain.NewCliente("cliente teste", "000.000.000-00", "cliente@teste.com.br")
	venda, _ := domain.NewVenda(cliente, time.Now(), "")

	categoria, _ := domain.NewCategoria("categoria teste")
	produto, _ := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")

	vendaItem, err := domain.NewVendaItem(venda, produto, 10)
	require.NotNil(t, vendaItem)
	require.Nil(t, err)
}
