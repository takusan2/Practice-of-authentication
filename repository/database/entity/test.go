package entity

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	ID   string `gorm:"type:varchar(36);primary_key;"`
	Name string `gorm:"type:varchar(255);not null;"`
}
