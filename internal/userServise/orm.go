package userService

import (
	"gorm.io/gorm"
	"pantela/internal/taskServise"
	"time"
)

type User struct {
	ID        string             `gorm:"primaryKey"`
	Email     string             `gorm:"unique;not null"`
	Password  string             `gorm:"not null"`
	DeletedAt gorm.DeletedAt     `gorm:"index"`
	CreatedAt time.Time          `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time          `gorm:"default:CURRENT_TIMESTAMP"`
	Tasks     []taskServise.Task `json:"tasks"`
}
