package dbconnect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

/*
SetDBConnection - Set Connection to Database
*/
func SetDBConnection() *gorm.DB {
	dsn := os.Getenv("DB_USER") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{QueryFields: true})

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(err)
		return nil
	}

	return db
}
