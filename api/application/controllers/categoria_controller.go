package controllers

import (
	"api/application/services"
	"net/http"
)

type CategoriaController interface {
	All(c Context) error
}

type categoriaController struct {
	service services.CategoriaService
}

func NewCategoriaController(service services.CategoriaService) CategoriaController {
	return &categoriaController{service}
}

func (controller *categoriaController) All(c Context) error {
	res, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
