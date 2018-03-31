package engine

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

func openDatabase() {
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(pwd)

	database, err = gorm.Open("sqlite3", pwd+"/template.db")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	database.DB().SetMaxOpenConns(100)
	database.DB().SetMaxIdleConns(1)
}

func init() {
	openDatabase()
}

//GetORM is
func GetORM() *gorm.DB {
	if database == nil {
		openDatabase()
	}
	return database.Debug()
}
