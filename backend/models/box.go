package models

import "gorm.io/gorm"

type Box struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Items      []Item `json:"items"`
}
