package database

import (
	"fmt"
	"os"

	sql "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/domain"
)

func Connect() *gorm.DB {
	// godotenv.Load("/Users/okadatakuya/my_folder/dev/my_app/（仮）/backend/.env")

	cfg := sql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  os.Getenv("NET"),
		Addr:                 os.Getenv("ADDR"),
		DBName:               os.Getenv("DBNAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	dsn := cfg.FormatDSN()

	db, err := gorm.Open("mysql", dsn)
	db.LogMode(true)
	if err != nil {
		fmt.Print("failed to Connected to database\n")
		panic(err.Error())
	}
	db.AutoMigrate(&domain.User{}, &domain.Note{})
	fmt.Print("Connected to database!!\n")
	return db
}
