package entity

import (
	abstraction "app-todo/models/base"
	"context"
	"time"

	"gorm.io/gorm"
)

type MasterCategoryEntity struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type MasterCategoryModel struct {
	MasterCategoryEntity
	abstraction.Entity
	Context context.Context `json:"-" gorm:"-"`
}

func (m *MasterCategoryModel) TableName() string {
	return "master_category"
}

func (m *MasterCategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.Created = time.Now()
	return
}

func (m *MasterCategoryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	update := time.Now()
	m.Updated = &update
	return
}

func (m *MasterCategoryModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
