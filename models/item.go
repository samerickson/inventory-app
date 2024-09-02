package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
}
