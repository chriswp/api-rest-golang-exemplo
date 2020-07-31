package router

import (
	"api/application/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/categorias", func(context echo.Context) error { return c.All(context) })
	e.POST("/categorias", func(context echo.Context) error { return c.All(context) })

	return e
}
