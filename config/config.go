package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var DbConfig Config

type Config interface {
	GetConnStr() string
}

type Sqlite struct {
	Path string
}

func (s Sqlite) GetConnStr() string {
	return s.Path
}

func newSqlite(v *viper.Viper) Sqlite {
	return Sqlite{Path: v.GetString("path")}
}

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v ", err))
	}

	DbConfig = getConfig(viper.GetViper())
}

func getConfig(v *viper.Viper) Config {
	var db Config
	switch v.Get("db.type") {
	case "sqlite":
		db = newSqlite(v.Sub("db.sqlite"))
	}

	return db
}
