package model

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Name        string
	Born        string
	Died        string
	Nationality string
	Wikipedia   string
	Photos      []Photo
}
