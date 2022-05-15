package config

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

var db *gorm.DB
var err error

func Database(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	autoCreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName)

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err = gorm.Open(Dbdriver, conn)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
}

func autoCreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/", DbUser, DbPassword, DbHost, DbPort)
	db, err := sql.Open(Dbdriver, dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DbName)
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return db
}
