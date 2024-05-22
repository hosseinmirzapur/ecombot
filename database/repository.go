package database

import (
	"fmt"

	"gorm.io/gorm"
)

func Create(model interface{}) *gorm.DB {
	return db.Create(model)
}

func FindOne(model interface{}, query string) *gorm.DB {
	return FindMany(model, query, 1)
}

func FindMany(model interface{}, clause string, take int) *gorm.DB {
	return db.Where(clause).Find(model).Limit(take)
}

func Update(model interface{}, searchKey string, searchValue interface{}, values map[string]interface{}) (*gorm.DB, error) {
	err := db.First(model, fmt.Sprintf("%s = ?", searchKey), searchValue).Error
	if err != nil {
		return nil, err
	}

	return db.Model(model).Updates(values), nil
}

func UpdateByID(model interface{}, id uint, values map[string]interface{}) (*gorm.DB, error) {
	return Update(model, "ID", id, values)
}

func Delete(model interface{}, searchKey string, searchValue interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", searchKey), searchValue).Delete(model)
}

func DeleteByID(model interface{}, modelID uint) *gorm.DB {
	return Delete(model, "ID", modelID)
}
