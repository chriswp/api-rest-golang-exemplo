package controllers

import (
	"api/application/services"
	"github.com/labstack/echo"
	"net/http"
)

type ProdutoController interface {
	All(c echo.Context) error
}

type produtoController struct {
	service services.ProdutoService
}

func NewProdutoController(service services.ProdutoService) ProdutoController {
	return &produtoController{service}
}

func (controller *produtoController) All(c echo.Context) error {
	res, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *produtoController) Show(c echo.Context) error {
	res, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
