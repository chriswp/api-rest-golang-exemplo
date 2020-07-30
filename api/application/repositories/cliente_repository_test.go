package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClienteRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	cliente, _ := domain.NewCliente("cliente", "000.000.000-00", "cliente@email.com.br")
	repo := repositories.NewClienteRepository(db)
	repo.Insert(cliente)

	c, err := repo.Find(cliente.ID)

	require.NotEmpty(t, c)
	require.Nil(t, err)
	require.Equal(t, c.ID, cliente.ID)
}
