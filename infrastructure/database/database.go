package database

import (
	"fmt"
	"os"

	sql "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/domain"
)

func Connect() *gorm.DB {

	cfg := sql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       os.Getenv("NET"),
		Addr:      os.Getenv("ADDR"),
		DBName:    os.Getenv("DBNAME"),
		ParseTime: true,
	}

	db, err := gorm.Open("mysql", cfg.FormatDSN())
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&domain.User{}, &domain.Note{})
	fmt.Print("Connected to database!!\n")
	return db
}
