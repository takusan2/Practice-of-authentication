package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/heart-note/domain"
	_ "gorm.io/driver/mysql"
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
	godotenv.Load("/Users/okadatakuya/my_folder/dev/my_app/（仮）/backend/.env")

	dsn := os.Getenv("CLEARDB_DATABASE_URL")
	fmt.Print("dsn", dsn)
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
