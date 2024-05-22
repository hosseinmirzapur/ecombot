package models

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	Name     string
	Products []*Product `gorm:"many2many:colors_products"`
}
