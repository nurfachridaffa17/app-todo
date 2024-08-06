package repository

import (
	"app-todo/models/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) Create(c echo.Context, DB *gorm.DB, category entity.TTaskCategoryModel) (entity.TTaskCategoryModel, error) {
	err := DB.Model(entity.TTaskCategoryModel{}).Create(&category).Error
	return category, err
}

func (r *repositoryImpl) Update(c echo.Context, DB *gorm.DB, TaskID int, categ entity.TTaskCategoryModel) (entity.TTaskCategoryModel, error) {
	existingRecord := entity.TTaskCategoryModel{}
	if err := DB.First(&existingRecord, "task_id = ?", TaskID).Error; err != nil {
		return existingRecord, err
	}
	updated_column := map[string]interface{}{
		"category_id": categ.CategoryID,
	}

	if err := DB.Model(&existingRecord).Updates(updated_column).Error; err != nil {
		return existingRecord, err
	}
	return existingRecord, nil
}
