package service

import (
	categ "app-todo/internal/category/repository"
	taskCateg "app-todo/internal/task-category/repository"
	"app-todo/internal/task/repository"
	"app-todo/models/base"
	"app-todo/models/dto"
	"app-todo/models/entity"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB           *gorm.DB
	Repository   repository.Repository
	TaskCategory taskCateg.Repository
	Category     categ.Repository
}

func NewService(
	DB *gorm.DB,
	Repository repository.Repository,
	TaskCategory taskCateg.Repository,
	Category categ.Repository,
) Service {
	return &serviceImpl{
		DB:           DB,
		Repository:   Repository,
		TaskCategory: TaskCategory,
		Category:     Category,
	}
}

func (s *serviceImpl) Create(c echo.Context, task dto.CreateTask) (response dto.CreateTaskRes, err error) {
	var data entity.MasterTaskModel

	dueDate, err := time.Parse("2006-01-02 15:04:05", task.DueDate)
	if err != nil {
		return response, fmt.Errorf("invalid date format: %v", err)
	}

	data, err = s.Repository.Create(c, s.DB, entity.MasterTaskModel{
		Entity: base.Entity{
			Createdby: task.Createdby,
		},
		MasterTaskEntity: entity.MasterTaskEntity{
			Title:       task.Title,
			Description: task.Description,
			DueDate:     dueDate,
			Status:      task.Status,
		},
	})

	if err != nil {
		return response, err
	}

	_, err = s.TaskCategory.Create(c, s.DB, entity.TTaskCategoryModel{
		Entity: base.Entity{},
		TTaskCategoryEntity: entity.TTaskCategoryEntity{
			TaskID:     data.ID,
			CategoryID: task.CategoryID,
		},
	})

	if err != nil {
		return response, err
	}

	nameCategory, err := s.Category.FindByID(c, s.DB, task.CategoryID)

	if err != nil {
		return response, err
	}

	response = dto.CreateTaskRes{
		MasterTaskModel: data,
		Category: dto.CategoryRes{
			MasterCategoryModel: nameCategory,
		},
	}

	return response, err
}

func (s *serviceImpl) Update(c echo.Context, ID int, task dto.UpdateTask) (response dto.CreateTaskRes, err error) {
	var data entity.MasterTaskModel

	dueDate, err := time.Parse("2006-01-02 15:04:05", task.DueDate)
	if err != nil {
		return response, fmt.Errorf("invalid date format: %v", err)
	}

	data, err = s.Repository.Update(c, s.DB, ID, entity.MasterTaskModel{
		Entity: base.Entity{
			Updatedby: &task.Updatedby,
		},
		MasterTaskEntity: entity.MasterTaskEntity{
			Title:       task.Title,
			Description: task.Description,
			DueDate:     dueDate,
			Status:      task.Status,
		},
	})

	if err != nil {
		return response, err
	}

	_, err = s.TaskCategory.Update(c, s.DB, ID, entity.TTaskCategoryModel{
		Entity: base.Entity{},
		TTaskCategoryEntity: entity.TTaskCategoryEntity{
			CategoryID: task.CategoryID,
		},
	})

	if err != nil {
		return response, err
	}

	nameCategory, err := s.Category.FindByID(c, s.DB, task.CategoryID)

	if err != nil {
		return response, err
	}

	response = dto.CreateTaskRes{
		MasterTaskModel: data,
		Category: dto.CategoryRes{
			MasterCategoryModel: nameCategory,
		},
	}

	return response, nil
}
