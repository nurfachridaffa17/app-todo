package repository

import (
	"app-todo/models/dto"
	"app-todo/models/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	// TODO: implement
	FindAll(c echo.Context, DB *gorm.DB, payload dto.GetCategoryRequest) (category []entity.MasterCategoryModel, err error, total int64)
	FindByID(c echo.Context, DB *gorm.DB, ID int) (category entity.MasterCategoryModel, err error)
	Create(c echo.Context, DB *gorm.DB, category entity.MasterCategoryModel) (entity.MasterCategoryModel, error)
	Update(c echo.Context, DB *gorm.DB, ID int, category entity.MasterCategoryModel) (entity.MasterCategoryModel, error)
	Delete(c echo.Context, DB *gorm.DB, id int) (entity.MasterCategoryModel, error)
}
