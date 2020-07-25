package repositories_test

import (
	"api/domain"
	"api/framework/database"
	"api/repositories"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestVendaRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	cliente, _ := domain.NewCliente("cliente teste", "000.000.000-00", "cliente@teste.com.br")
	venda, _ := domain.NewVenda(cliente, time.Now(), "")
	repo := repositories.NewVendaRepository(db)
	repo.Insert(venda)

	v, err := repo.Find(venda.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, venda.ID)
}
