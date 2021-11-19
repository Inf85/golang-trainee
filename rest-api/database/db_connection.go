package dbconnect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
SetDBConnection - Set DataBase Connection
 */
func SetDBConnection() *gorm.DB {
	dsn:= "root@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ QueryFields: true})


	if err != nil{
		panic(err)
		return nil
	}

	return db
}