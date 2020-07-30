package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCompraRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	compra, _ := domain.NewCompra(time.Now())
	categoria, _ := domain.NewCategoria("categoria 1")
	produto, _ := domain.NewProduto(categoria, "produto 1", 100.00, "apenas para efeito de teste")
	produto2, _ := domain.NewProduto(categoria, "produto 2", 100.00, "apenas para efeito de teste")
	compra.AddNewItem(produto, 2)
	compra.AddNewItem(produto2, 1)

	repo := repositories.NewCompraRepository(db)
	repo.Insert(compra)

	c, err := repo.Find(compra.ID)

	require.NotEmpty(t, c.ID)
	require.Equal(t, len(compra.Itens), len(c.Itens))
	require.Nil(t, err)
	require.Equal(t, c.ID, compra.ID)
}
