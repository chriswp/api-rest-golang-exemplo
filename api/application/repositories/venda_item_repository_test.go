package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestVendaItemRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	cliente, _ := domain.NewCliente("cliente teste", "000.000.000-00", "cliente@teste.com.br")
	venda, _ := domain.NewVenda(cliente, time.Now(), "")

	categoria, _ := domain.NewCategoria("categoria teste")
	produto, err := domain.NewProduto(categoria, "produto", 10.50, "apenas para teste")

	item, _ := domain.NewVendaItem(venda, produto, 2)
	fmt.Println("sete", item.ID)
	repo := repositories.NewVendaItemRepository(db)
	repo.Insert(item)

	vi, err := repo.Find(item.ID)

	require.NotEmpty(t, vi.ID)
	require.Nil(t, err)
	require.Equal(t, vi.ID, item.ID)
}
