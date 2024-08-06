package repository

import (
	"app-todo/models/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	Create(c echo.Context, DB *gorm.DB, category entity.TTaskCategoryModel) (entity.TTaskCategoryModel, error)
	Update(c echo.Context, DB *gorm.DB, TaskID int, categ entity.TTaskCategoryModel) (entity.TTaskCategoryModel, error)
}
