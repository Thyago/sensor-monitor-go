package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	driver string
	url    string
	ormDB  *gorm.DB
}

func NewDatabase(user, password, host, port, name string) *Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	return &Database{driver: "mysql", url: url}
}

func (db *Database) Open() error {
	gDB, err := gorm.Open(mysql.Open(db.url), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database: %v", db.driver, err)
		return err
	}
	db.ormDB = gDB
	return nil
}

func (db *Database) Migrate(values ...interface{}) {
	db.ormDB.Debug().AutoMigrate(values...)
}

func (db *Database) GetORM() *gorm.DB {
	return db.ormDB
}
