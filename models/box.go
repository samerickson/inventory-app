package models

import "gorm.io/gorm"

type Box struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
}
