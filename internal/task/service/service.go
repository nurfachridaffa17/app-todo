package service

import (
	"app-todo/models/dto"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Create(c echo.Context, task dto.CreateTask) (dto.CreateTaskRes, error)
	Update(c echo.Context, ID int, task dto.UpdateTask) (dto.CreateTaskRes, error)
}
