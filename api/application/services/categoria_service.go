package services

import (
	"api/application/presenters"
	"api/application/repositories"
	"api/domain"
)

type CategoriaService interface {
	Get(id string) (*domain.Categoria, error)
	GetAll(categorias []*domain.Categoria) ([]*domain.Categoria, error)
	Create(categoria *domain.Categoria) (*domain.Categoria, error)
	Edit(categoria *domain.Categoria, id string) (*domain.Categoria, error)
}

type categoriaService struct {
	Repository repositories.CategoriaRepository
	Presenter  presenters.CategoriaPresenter
}

func NewCategoriaService(r repositories.CategoriaRepository, p presenters.CategoriaPresenter) CategoriaService {
	return &categoriaService{r, p}
}

func (cs *categoriaService) Get(id string) (*domain.Categoria, error) {
	c, err := cs.Repository.Find(id)
	if err != nil {
		return nil, err
	}

	return cs.Presenter.Response(c), nil
}

func (cs *categoriaService) GetAll(categorias []*domain.Categoria) ([]*domain.Categoria, error) {
	c, err := cs.Repository.All(categorias)
	if err != nil {
		return nil, err
	}

	return cs.Presenter.ResponseAll(c), nil
}

func (cs *categoriaService) Create(categoria *domain.Categoria) (*domain.Categoria, error) {
	c, err := cs.Repository.Insert(categoria)
	if err != nil {
		return nil, err
	}

	return cs.Presenter.Response(c), nil
}

func (cs *categoriaService) Edit(categoria *domain.Categoria, id string) (*domain.Categoria, error) {
	c, err := cs.Repository.Update(categoria, id)
	if err != nil {
		return nil, err
	}

	return cs.Presenter.Response(c), nil
}
