package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewVenda(t *testing.T) {
	cliente, _ := domain.NewCliente("cliente teste", "000.000.000-00", "cliente@teste.com.br")
	venda, err := domain.NewVenda(cliente, time.Now(), "")
	require.NotNil(t, venda)
	require.Nil(t, err)
}
