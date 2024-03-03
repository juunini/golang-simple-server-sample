package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement;not null;"`
	Title     string         `json:"title" gorm:"type:varchar(50);not null;"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
