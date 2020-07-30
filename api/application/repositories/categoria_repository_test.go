package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCategoriaRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	categoria, _ := domain.NewCategoria("categoria")
	repo := repositories.NewCategoriaRepository(db)
	repo.Insert(categoria)

	c, err := repo.Find(categoria.ID)

	require.NotEmpty(t, c.ID)
	require.Nil(t, err)
	require.Equal(t, c.ID, categoria.ID)
}
