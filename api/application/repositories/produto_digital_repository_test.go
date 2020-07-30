package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProdutoDigitalRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	categoria, _ := domain.NewCategoria("categoria teste")
	produto, _ := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")

	pDigital, _ := domain.NewProdutoDigital(produto, "http://teste.com.br")
	repo := repositories.NewProdutoDigitalRepository(db)
	repo.Insert(pDigital)

	pd, err := repo.Find(pDigital.ID)

	require.NotEmpty(t, pd.ID)
	require.Nil(t, err)
	require.Equal(t, pd.ID, pDigital.ID)
}
