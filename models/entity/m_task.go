package entity

import (
	abstraction "app-todo/models/base"
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type MasterTaskEntity struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type MasterTaskModel struct {
	MasterTaskEntity
	abstraction.Entity
	Context context.Context `json:"-" gorm:"-"`
}

func (m MasterTaskModel) MarshalJSON() ([]byte, error) {
	type Alias MasterTaskModel
	return json.Marshal(&struct {
		DueDate string `json:"due_date"`
		*Alias
	}{
		DueDate: m.DueDate.Format("2006-01-02 15:04:05"), // Custom date format
		Alias:   (*Alias)(&m),
	})
}

func (m *MasterTaskModel) TableName() string {
	return "master_tasks"
}

func (m *MasterTaskModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.Created = time.Now()
	return
}

func (m *MasterTaskModel) BeforeUpdate(tx *gorm.DB) (err error) {
	update := time.Now()
	m.Updated = &update
	return
}

func (m *MasterTaskModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
