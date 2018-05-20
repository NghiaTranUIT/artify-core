package model

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Name     string `json:"born"`
	ImageUrl string `gorm:"not null" json:"image_url"`
	Author   Author `json:"author"`
	AuthorID uint   `json:"-"`
	Width    uint   `gorm:"not null" json:"width"`
	Height   uint   `gorm:"not null" json:"height"`
}
