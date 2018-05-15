package model

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Name     string
	ImageUrl *string `gorm:"not null"`
	AuthorID uint
}
