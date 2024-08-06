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

func (r *repositoryImpl) Create(c echo.Context, DB *gorm.DB, task entity.MasterTaskModel) (entity.MasterTaskModel, error) {
	err := DB.Model(entity.MasterTaskModel{}).Create(&task).Error
	return task, err
}

func (r *repositoryImpl) Update(c echo.Context, DB *gorm.DB, ID int, task entity.MasterTaskModel) (entity.MasterTaskModel, error) {
	existingRecord := entity.MasterTaskModel{}
	if err := DB.First(&existingRecord, "id = ?", ID).Error; err != nil {
		return existingRecord, err
	}
	updated_column := map[string]interface{}{
		"name":        task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
		"status":      task.Status,
		"updatedby":   task.Updatedby,
	}

	if err := DB.Model(&existingRecord).Updates(updated_column).Error; err != nil {
		return existingRecord, err
	}
	return existingRecord, nil
}
