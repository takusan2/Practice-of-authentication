package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/domain"
)

func Connect() *gorm.DB {

	// cfg := sql.Config{
	// 	User:                 os.Getenv("DBUSER"),
	// 	Passwd:               os.Getenv("DBPASS"),
	// 	Net:                  os.Getenv("NET"),
	// 	Addr:                 os.Getenv("ADDR"),
	// 	DBName:               os.Getenv("DBNAME"),
	// 	ParseTime:            true,
	// 	AllowNativePasswords: true,
	// }
	// dsn := cfg.FormatDSN()
	dsn := os.Getenv("CLEARDB_DATABASE_URL")
	db, err := gorm.Open("mysql", dsn)
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&domain.User{}, &domain.Note{})
	fmt.Print("Connected to database!!\n")
	return db
}
