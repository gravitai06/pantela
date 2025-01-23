package userService

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Исправлено на string
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
}

//package userService
//
//import "gorm.io/gorm"
//
//type User struct {
//	gorm.Model
//	Email    string `json:"email" gorm:"unique;not null"`
//	Password string `json:"password" gorm:"not null"`
//}
