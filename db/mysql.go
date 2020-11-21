package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbname = "mimuse_db"
	dbhost = "lp"
	dbuser = "root"
	dbpass = "12345678"
	dbport = "3306"
)

func Connect() *gorm.DB {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname))
	if err != nil {
		fmt.Errorf("Can't connect to database")
	}

	return db
}
