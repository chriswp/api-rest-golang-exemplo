package services

import (
	"api/application/presenters"
	"api/application/repositories"
	"api/domain"
)

type ProdutoService interface {
	GetAll() ([]*domain.Produto, error)
}

type produtoService struct {
	Repository repositories.ProdutoRepository
	Presenter  presenters.ProdutoPresenter
}

func NewProdutoService(r repositories.ProdutoRepository, p presenters.ProdutoPresenter) ProdutoService {
	return &produtoService{r, p}
}

func (ps *produtoService) GetAll() ([]*domain.Produto, error) {
	c, err := ps.Repository.All()
	if err != nil {
		return nil, err
	}

	return ps.Presenter.Response(c), nil
}
