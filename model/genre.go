package model

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name        string `gorm:"size:100" json:"name" validate:"required"`
	Description string `gorm:"type:text" json:"description"`

	Books []Book
}
