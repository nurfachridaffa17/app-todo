package service

import (
	"app-todo/models/dto"

	"github.com/labstack/echo/v4"
)

type Service interface {
	FindAll(c echo.Context, payload dto.GetCategoryRequest) (response []dto.CategoryRes, err error, total int64)
	FindByID(c echo.Context, id int) (response dto.CategoryRes, err error)
	Create(c echo.Context, category dto.CreateCategory) (response dto.CategoryRes, err error)
	Update(c echo.Context, id int, category dto.UpdateCategory) (response dto.CategoryRes, err error)
	Delete(c echo.Context, id int) (response dto.CategoryRes, err error)
}
