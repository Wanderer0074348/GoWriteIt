package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null" validate:"required"`
	Content string `json:"content" gorm:"not null" validate:"required"`
}
