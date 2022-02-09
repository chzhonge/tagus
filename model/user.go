package model

import "gorm.io/gorm"

type User struct {
	ID          uint   `gorm:"column:id;unsigned;autoIncrement;"`
	UserName    string `gorm:"column:username;type:varchar(50);unique;index;not null;"`
	Password    string `gorm:"column:password;type:char(60);not null;"`
	DisplayName string `gorm:"column:displayname;type:varchar(20);not null;"`
	gorm.Model
}
