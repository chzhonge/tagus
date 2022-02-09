package model

import "gorm.io/gorm"

type Term struct {
	ID      uint   `gorm:"primaryKey;column:id;type:int(10);unsigned;"`
	UserID  uint   `gorm:"column:user_id;type:int(10);unsigned;index;"`
	Name    string `gorm:"column:name;type:varchar(60);not null;index;"`
	Mapping string `gorm:"column:mapping;type:text;not null"`
	gorm.Model
}
