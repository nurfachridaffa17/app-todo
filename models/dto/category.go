package dto

import (
	model "app-todo/models/entity"
	res "app-todo/pkg/util/response"
)

type (
	GetCategoryRequest struct {
		Page   int    `queryparam:"page"`
		Limit  int    `queryparam:"limit"`
		Search string `queryparam:"search"`
		Filter int    `queryparam:"filter"`
		Name   string `queryparam:"name"`
		UserID string `queryparam:"user_id"`
	}

	CreateCategory struct {
		Name      string `json:"name"`
		UserID    uint   `json:"user_id"`
		CreatedBy int    `json:"created_by"`
	}

	UpdateCategory struct {
		Name      string `json:"name"`
		UserID    uint   `json:"user_id"`
		UpdatedBy int    `json:"updated_by"`
	}

	DeleteCategory struct {
		ID        int `json:"id"`
		DeletedBy int `json:"deleted_by"`
	}
)

type (
	CategoryRes struct {
		model.MasterCategoryModel
	}

	CategoryResponse struct {
		Body struct {
			Meta res.Meta    `json:"meta"`
			Data CategoryRes `json:"data"`
		} `json:"body"`
	}
)
