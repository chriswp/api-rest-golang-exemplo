package repositories_test

import (
	"api/application/repositories"
	"api/domain"
	"api/framework/database"
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

	categoria, _ := domain.NewCategoria()
	produto, _ := domain.NewProduto(categoria, "produto 1", 100.00, "apenas para efeito de teste")
	produto2, _ := domain.NewProduto(categoria, "produto 1", 100.00, "apenas para efeito de teste")

	venda.AddNewItem(produto, 1)
	venda.AddNewItem(produto2, 5)
	repo.Insert(venda)

	v, err := repo.Find(venda.ID)

	require.NotEmpty(t, v.ID)
	require.Equal(t, len(venda.Itens), len(v.Itens))
	require.Nil(t, err)
	require.Equal(t, v.ID, venda.ID)
}
