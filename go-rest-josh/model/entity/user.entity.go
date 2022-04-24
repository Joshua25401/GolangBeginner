package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" form:"id" gorm:"primaryKey" `
	Name      string         `json:"name" form:"id" validate:"required"`
	Email     string         `json:"email" form:"email" validate:"required,email"`
	Address   string         `json:"address" form:"address" validate:"required"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
