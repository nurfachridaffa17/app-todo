package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
