package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tagus/config"
)

var DBConn *gorm.DB

func Init() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open(config.DbConfig.GetConnStr()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error DBConn : %v ", err))
	}
}
