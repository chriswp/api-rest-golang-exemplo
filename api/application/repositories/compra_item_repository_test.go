package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCompraItemRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	compra, _ := domain.NewCompra(time.Now())

	categoria, _ := domain.NewCategoria("categoria teste")
	produto, _ := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")

	item, _ := domain.NewCompraItem(compra, produto, 2)
	repo := repositories.NewCompraItemRepository(db)
	repo.Insert(item)

	c, err := repo.Find(item.ID)

	require.NotEmpty(t, c.ID)
	require.Nil(t, err)
	require.Equal(t, c.ID, item.ID)
}
