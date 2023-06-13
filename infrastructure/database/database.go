package database

import (
	"fmt"
	"os"

	sql "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/domain"
)

func Connect() *gorm.DB {
	var dsn string
	cfg := sql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       os.Getenv("DB_NET"),
		Addr:      os.Getenv("DB_ADDR"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
	dsn = cfg.FormatDSN()
	fmt.Print(dsn)
	// dsn = os.Getenv("DATABASE_URL") + "/?parseTime=true"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Print("failed to Connected to database\n")
		panic(err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&domain.User{}, &domain.Note{})
	fmt.Print("Connected to database!!\n")
	return db
}
