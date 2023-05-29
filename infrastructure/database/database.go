package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/takuya-okada-01/heart-note/domain"
)

func Connect() *gorm.DB {
	var dsn string

	// cfg := sql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    os.Getenv("NET"),
	// 	Addr:   os.Getenv("ADDR"),
	// 	DBName: os.Getenv("DBNAME"),
	// }
	// dsn = cfg.FormatDSN()
	dsn = os.Getenv("DATABASE_URL")

	fmt.Print("Try connecting to database...\n")
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
