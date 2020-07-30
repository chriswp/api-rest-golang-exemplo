package services

import (
	"api/application/presenter"
	"api/application/repositories"
	"api/domain"
)

type ProdutoService interface {
	GetAll() ([]*domain.Produto, error)
}

type produtoService struct {
	Repository repositories.ProdutoRepository
	Presenter  presenter.ProdutoPresenter
}

func NewProdutoService(r repositories.ProdutoRepository, p presenter.ProdutoPresenter) ProdutoService {
	return &produtoService{r, p}
}

func (ps *produtoService) GetAll() ([]*domain.Produto, error) {
	c, err := ps.Repository.All()
	if err != nil {
		return nil, err
	}

	return ps.Presenter.Response(c), nil
}
