package presenters

import "api/domain"

type ProdutoPresenter interface {
	Response(produtos []*domain.Produto) []*domain.Produto
}

type produtoPresenter struct {
}

func NewProdutoPresenter() ProdutoPresenter {
	return &produtoPresenter{}
}

func (cp produtoPresenter) Response(produtos []*domain.Produto) []*domain.Produto {
	for _, p := range produtos {
		p.ID = p.ID
		p.Nome = p.Nome
		p.Categoria.Nome = p.Categoria.Nome
		p.Ativo = p.Ativo
	}
	return produtos
}
