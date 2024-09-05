package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	BoxId      uint   `json:"box"`
}
