package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := ""
	PROTOCOL := "tcp(localhost)"  //tcp(##.###.##.###:3306)
	DBNAME   := "GoApiServer"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	db.DB()

	DB = db
}
