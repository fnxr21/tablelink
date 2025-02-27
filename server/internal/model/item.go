package model

import "gorm.io/gorm"

type TmItem struct {
	gorm.Model
	Uuid  string
	Name  string
	Price float64 `gorm:"type:decimal(15,2)"`
}
