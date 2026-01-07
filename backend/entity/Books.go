package entity

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title string  `valid:"required,stringlength(3|100)"`
	Price float64 `valid:"required,range(50|5000)"`
	Code  string  `valid:"required,^^BK//d[6]"`
}
