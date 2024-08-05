package service

import (
	"app-todo/internal/category/repository"
	"app-todo/models/base"
	"app-todo/models/dto"
	"app-todo/models/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB         *gorm.DB
	Repository repository.Repository
}

func NewService(
	DB *gorm.DB,
	Repository repository.Repository,
) Service {
	return &serviceImpl{
		DB:         DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) FindAll(c echo.Context, payload dto.GetCategoryRequest) (response []dto.CategoryRes, err error, total int64) {
	dataList, err, total := s.Repository.FindAll(c, s.DB, payload)
	if err != nil {
		return response, err, total
	}

	for _, v := range dataList {
		response = append(response, dto.CategoryRes{
			MasterCategoryModel: v,
		})
	}

	return response, err, total
}

func (s *serviceImpl) FindByID(c echo.Context, id int) (response dto.CategoryRes, err error) {
	data, err := s.Repository.FindByID(c, s.DB, id)
	if err != nil {
		return response, err
	}
	response = dto.CategoryRes{
		MasterCategoryModel: data,
	}

	return response, err
}

func (s *serviceImpl) Create(c echo.Context, category dto.CreateCategory) (response dto.CategoryRes, err error) {
	var data entity.MasterCategoryModel

	data, err = s.Repository.Create(c, s.DB, entity.MasterCategoryModel{
		Entity: base.Entity{
			Createdby: category.CreatedBy,
		},
		MasterCategoryEntity: entity.MasterCategoryEntity{
			Name:   category.Name,
			UserID: category.UserID,
		},
	})

	if err != nil {
		return response, err
	}

	response = dto.CategoryRes{
		MasterCategoryModel: data,
	}

	return response, err
}

func (s *serviceImpl) Update(c echo.Context, id int, payload dto.UpdateCategory) (response dto.CategoryRes, err error) {
	data, err := s.Repository.Update(c, s.DB, id, entity.MasterCategoryModel{
		Entity: base.Entity{
			Updatedby: &payload.UpdatedBy,
		},
		MasterCategoryEntity: entity.MasterCategoryEntity{
			Name:   payload.Name,
			UserID: payload.UserID,
		},
	})

	if err != nil {
		return response, err
	}

	response = dto.CategoryRes{
		MasterCategoryModel: data,
	}

	return response, nil
}

func (s *serviceImpl) Delete(c echo.Context, id int) (response dto.CategoryRes, err error) {
	data, err := s.Repository.Delete(c, s.DB, id)
	if err != nil {
		return response, err
	}

	response = dto.CategoryRes{
		MasterCategoryModel: data,
	}

	return response, err
}
