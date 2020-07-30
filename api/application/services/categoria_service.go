package services

import (
	"api/application/presenter"
	"api/application/repositories"
	"api/domain"
)

type CategoriaService interface {
	GetAll() ([]*domain.Categoria, error)
}

type categoriaService struct {
	Repository repositories.CategoriaRepository
	Presenter  presenter.CategoriaPresenter
}

func NewCategoriaService(r repositories.CategoriaRepository, p presenter.CategoriaPresenter) *categoriaService {
	return &categoriaService{r, p}
}

func (cs *categoriaService) GetAll() ([]*domain.Categoria, error) {
	c, err := cs.Repository.All()
	if err != nil {
		return nil, err
	}

	return cs.Presenter.Response(c), nil
}
