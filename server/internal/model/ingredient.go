package model

import "gorm.io/gorm"

type TmIngredient struct {
	gorm.Model
	Uuid        string
	Name        string
	CauseAlergy string
	Type        int
	Status      int
}
