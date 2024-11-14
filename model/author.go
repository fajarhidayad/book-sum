package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `gorm:"size:100" json:"name" validate:"required"`
	Bio  string `gorm:"type:text" json:"bio"`

	Books []Book
}
