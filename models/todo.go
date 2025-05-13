package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo 任务模型
type Todo struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:100;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	DueDate     *time.Time     `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
