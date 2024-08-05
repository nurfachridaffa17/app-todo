package entity

import (
	abstraction "app-todo/models/base"
	"context"
)

type TTaskCategoryEntity struct {
	TaskID     int `json:"task_id" gorm:"primaryKey"`
	CategoryID int `json:"category_id" gorm:"primaryKey"`
}

type TTaskCategoryModel struct {
	TTaskCategoryEntity
	abstraction.Entity
	Context context.Context `json:"-" gorm:"-"`
}

func (m *TTaskCategoryModel) TableName() string {
	return "t_task_category"
}
