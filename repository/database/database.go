package database

import (
	"fmt"
	"os"

	sql "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
)

func Connect() *gorm.DB {
	err := godotenv.Load("/Users/okadatakuya/my_folder/dev/my_app/（仮）/backend/.env")
	if err != nil {
		fmt.Println(err.Error())
	}

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
	db.AutoMigrate(&entity.User{})
	fmt.Print("Connected to database!!\n")
	return db
}
