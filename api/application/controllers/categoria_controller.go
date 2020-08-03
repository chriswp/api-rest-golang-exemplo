package controllers

import (
	"api/application/services"
	"api/domain"
	"github.com/labstack/echo"
	"net/http"
)

type CategoriaController interface {
	All(c echo.Context) error
	Store(c echo.Context) error
	Update(c echo.Context) error
	Show(c echo.Context) error
}

type categoriaController struct {
	service services.CategoriaService
}

func NewCategoriaController(service services.CategoriaService) CategoriaController {
	return &categoriaController{service}
}

func (controller *categoriaController) All(c echo.Context) error {
	var categorias []*domain.Categoria
	res, err := controller.service.GetAll(categorias)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *categoriaController) Store(c echo.Context) error {
	categoria := domain.Categoria{}
	if err := c.Bind(&categoria); err != nil {
		return err
	}

	res, err := controller.service.Create(&categoria)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (controller *categoriaController) Update(c echo.Context) error {
	categoria := domain.Categoria{}
	id := c.Param("id")
	if err := c.Bind(&categoria); err != nil {
		return err
	}

	res, err := controller.service.Edit(&categoria, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *categoriaController) Show(c echo.Context) error {
	id := c.Param("id")
	res, err := controller.service.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
