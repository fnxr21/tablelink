package model

import "gorm.io/gorm"

type TmItem struct {
	gorm.Model
	Uuid  string
	Name  string
	Price float64 `gorm:"type:decimal(15,2)"`
}
type TmIngredient struct {
	gorm.Model
	Uuid        string
	Name        string
	CauseAlergy string
	Type        int
	Status      int
}
type TmItemIngredient struct {
	gorm.Model
	UuuidItem        string
	UuuidItems       TmItem `gorm:"references:Uuid"`
	UuuidIngredient  string
	UuuidIngredients TmIngredient `gorm:"references:Uuid"`
}
