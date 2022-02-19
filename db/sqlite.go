package db

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

type MockDialector struct {
	sqlite.Dialector
}

func (dialector MockDialector) Initialize(db *gorm.DB) (err error) {
	if dialector.DriverName == "" {
		dialector.DriverName = sqlite.DriverName
	}

	if dialector.Conn != nil {
		db.ConnPool = dialector.Conn
	} else {
		conn, err := sql.Open(dialector.DriverName, dialector.DSN)
		if err != nil {
			return err
		}
		db.ConnPool = conn
	}
	
	callbacks.RegisterDefaultCallbacks(db, &callbacks.Config{
		LastInsertIDReversed: true,
	})

	for k, v := range dialector.ClauseBuilders() {
		db.ClauseBuilders[k] = v
	}
	return
}
