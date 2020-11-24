package db

import (
	"database/sql"
	"fmt"

	/*Importing mysql driver */
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbuser = "root"
	dbpass = "password"
	//dbhost = "127.0.0.1"
	dbhost = "mysql"
	dbport = "3306"
	dbname = "mimuse_db"
)

//CreateCon open a mysql connection
func Connect() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname))
	if err != nil {
		log.Println("MySQL db is not connected")
		log.Println(err.Error())
		log.Fatal(err)
	}
	log.Println("MySQL connected")
	return db
}
