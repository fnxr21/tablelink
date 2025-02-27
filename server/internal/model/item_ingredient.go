package model

import "gorm.io/gorm"

type TmItemIngredient struct {
	gorm.Model
	UuuidItem        string
	UuuidItems       TmItem `gorm:"references:Uuid"`
	UuuidIngredient  string
	UuuidIngredients TmIngredient `gorm:"references:Uuid"`
}
