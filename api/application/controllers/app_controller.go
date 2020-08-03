package controllers

import "github.com/labstack/echo"

type AppController interface {
	All(c echo.Context) error
	Store(c echo.Context) error
	Update(c echo.Context) error
	Show(c echo.Context) error
}
