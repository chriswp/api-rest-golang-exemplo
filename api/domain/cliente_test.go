package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCliente(t *testing.T) {
	cliente, err := domain.NewCliente("cliente teste", "000.000.000-00", "cliente@teste.com.br")
	require.NotNil(t, cliente)
	require.Nil(t, err)
}

func TestNewClienteEmailInvalid(t *testing.T) {
	_, err := domain.NewCliente("cliente teste", "000.000.000-00", "email")
	require.Error(t, err)
}

func TestNewClienteFieldsEmpty(t *testing.T) {
	_, err := domain.NewCliente("", "", "")
	require.Error(t, err)

	_, nomeEmpty := domain.NewCliente("", "000.000.000-00", "cliente@teste.com.br")
	require.Error(t, nomeEmpty)

	_, cpfEmpty := domain.NewCliente("cliente", "", "cliente@teste.com.br")
	require.Error(t, cpfEmpty)

	_, emailEmpty := domain.NewCliente("cliente", "000.000.000-00", "")
	require.Error(t, emailEmpty)
}
