package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	URL     string
	Product Product

	ProductID uint
}
