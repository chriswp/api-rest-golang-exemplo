package repositories_test

import (
	"api/domain"
	"api/framework/database"
	"api/repositories"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCompraRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	compra, _ := domain.NewCompra(time.Now())
	repo := repositories.NewCompraRepository(db)
	repo.Insert(compra)

	c, err := repo.Find(compra.ID)

	require.NotEmpty(t, c.ID)
	require.Nil(t, err)
	require.Equal(t, c.ID, compra.ID)
}
