package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProdutoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	categoria, _ := domain.NewCategoria("categoria teste")
	produto, err := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")
	repo := repositories.NewProdutoRepository(db)
	repo.Insert(produto)

	c, err := repo.Find(produto.ID)

	require.NotEmpty(t, c.ID)
	require.Nil(t, err)
	require.Equal(t, c.ID, produto.ID)
}
