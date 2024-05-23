package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	URL       string
	Product   Product
	ProductID uint
}
