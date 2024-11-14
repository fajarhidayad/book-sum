package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `gorm:"size:255" json:"title" validate:"required"`
	PageCount int    `json:"page_count" validate:"required,gte=10"`
	AuthorID  uint   `json:"author_id" validate:"required"`
	Author    Author
	GenreID   uint `json:"genre_id" validate:"required"`
	Genre     Genre
}
