package dto

import (
	model "app-todo/models/entity"
	res "app-todo/pkg/util/response"
)

type (
	CreateTaskCategory struct {
		TaskID     int `json:"task_id"`
		CategoryID int `json:"category_id"`
	}

	UpdateTaskCategory struct {
		TaskID     int `json:"task_id"`
		CategoryID int `json:"category_id"`
	}

	DeleteTaskCategory struct {
		TaskID     int `json:"task_id"`
		CategoryID int `json:"category_id"`
	}
)

type (
	TaskCategoryRes struct {
		model.TTaskCategoryModel
	}

	TaskCategoryResponse struct {
		Body struct {
			Meta res.Meta        `json:"meta"`
			Data TaskCategoryRes `json:"data"`
		} `json:"body"`
	}
)
