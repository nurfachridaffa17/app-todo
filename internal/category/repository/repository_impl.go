package repository

import (
	"app-todo/models/dto"
	"app-todo/models/entity"
	"app-todo/pkg/paginate"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) FindAll(c echo.Context, DB *gorm.DB, payload dto.GetCategoryRequest) (category []entity.MasterCategoryModel, err error, total int64) {
	query := DB.Order("id DESC")
	payload.Search = strings.ToLower(payload.Search)
	wildcard := "%" + payload.Search + "%"

	if payload.Search != "" {
		query = query.Where("lower(name) like ?", wildcard)
	} else if payload.Filter != 0 {
		if payload.Name != "" {
			query = query.Where("name = ?", payload.Name)
		} else if payload.UserID != "" {
			query = query.Where("user_id = ?", payload.UserID)
		}
	}

	query.Model(entity.MasterCategoryModel{}).Count(&total)

	if payload.Page == 0 {
		err = query.Find(&category).Error
	} else {
		paginator := paginate.NewPaginate(payload.Page, payload.Limit)
		err = query.Scopes(paginator.PaginatedResult).Find(&category).Error
	}

	return category, err, total
}

func (r *repositoryImpl) FindByID(c echo.Context, DB *gorm.DB, ID int) (category entity.MasterCategoryModel, err error) {
	err = DB.Model(entity.MasterCategoryModel{}).Where("id = ?", ID).First(&category).Error
	return category, err
}

func (r *repositoryImpl) Create(c echo.Context, DB *gorm.DB, category entity.MasterCategoryModel) (entity.MasterCategoryModel, error) {
	err := DB.Model(entity.MasterCategoryModel{}).Create(&category).Error
	return category, err
}

func (r *repositoryImpl) Update(c echo.Context, DB *gorm.DB, ID int, categ entity.MasterCategoryModel) (entity.MasterCategoryModel, error) {
	existingRecord := entity.MasterCategoryModel{}
	if err := DB.First(&existingRecord, "id = ?", ID).Error; err != nil {
		return existingRecord, err
	}

	updated_column := map[string]interface{}{
		"name":      categ.Name,
		"user_id":   categ.UserID,
		"updatedby": categ.Updatedby,
	}

	if err := DB.Model(&existingRecord).Updates(updated_column).Error; err != nil {
		return existingRecord, err
	}

	return existingRecord, nil
}

func (r *repositoryImpl) Delete(c echo.Context, DB *gorm.DB, id int) (entity.MasterCategoryModel, error) {
	existingRecord := entity.MasterCategoryModel{}
	if err := DB.Model(entity.MasterCategoryModel{}).Where("id = ?", id).First(&existingRecord).Error; err != nil {
		return existingRecord, err
	}

	if err := DB.Model(entity.MasterCategoryModel{}).Where("id = ?", id).Delete(&existingRecord).Error; err != nil {
		return existingRecord, err
	}
	return existingRecord, nil
}
