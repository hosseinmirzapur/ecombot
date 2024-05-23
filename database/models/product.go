package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string
	Description string
	Price       string
	Code        string
	Videos      []*Video
	Images      []*Image
	Colors      []*Color `gorm:"many2many:colors_products"`
}
