package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //+ Ease 2018.11.23 改为sqlite3数据库
	//_ "github.com/Go-SQL-Driver/MySQL" //- Ease 2018.11.23 改为sqlite3数据库
	//"strconv" //- Ease 2018.11.23 改为sqlite3数据库
)

var db *gorm.DB

func InitDB() {
	err := connectDB()
	if err != nil {
		panic(err)
	}

	if (!db.HasTable(&Host{})) {
		db.CreateTable(&Host{})
	}

	if (!db.HasTable(&Page{})) {
		db.CreateTable(&Page{})
	}

	if (!db.HasTable(&DailyRecord{})) {
		db.CreateTable(&DailyRecord{})
	}

	if (!db.HasTable(&MonthlyRecord{})) {
		db.CreateTable(&MonthlyRecord{})
	}

}

func connectDB() error {
	config := GetConfig()
	var err error

	//db, err = gorm.Open("mysql", config.DB.Username+":"+config.DB.Password+
	//	"@tcp("+config.DB.Host+":"+strconv.FormatInt(config.DB.Port, 10)+")/"+
	//	config.DB.DBName+"?charset=utf8&parseTime=True")

	db, err = gorm.Open("sqlite3", config.DB.DBName)
	return err
}

func GetDB() *gorm.DB {
	if db == nil {
		err := connectDB()
		if err != nil {
			panic(err)
		}
	}
	return db
}
