package controllers

import (
	"api/application/services"
	"net/http"
)

type ProdutoController interface {
	All(c Context) error
}

type produtoController struct {
	service services.ProdutoService
}

func NewProdutoController(service services.ProdutoService) ProdutoController {
	return &produtoController{service}
}

func (controller *produtoController) All(c Context) error {
	res, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *produtoController) Show(c Context) error {
	res, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
