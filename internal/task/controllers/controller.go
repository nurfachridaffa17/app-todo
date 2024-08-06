package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
}
