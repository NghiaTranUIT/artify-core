package model

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Name        string  `json:"name"`
	Born        string  `json:"born"`
	Died        string  `json:"died"`
	Nationality string  `json:"nationality"`
	Wikipedia   string  `json:"wikipedia"`
	Photos      []Photo `json:"photos"`
}
