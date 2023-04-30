package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func Connect() *bun.DB {
	err := godotenv.Load("/Users/okadatakuya/my_folder/dev/my_app/（仮）/backend/.env")
	if err != nil {
		fmt.Println(err.Error())
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    os.Getenv("NET"),
		Addr:   os.Getenv("ADDR"),
		DBName: os.Getenv("DBNAME"),
	}
	sqldb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	if err := sqldb.Ping(); err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	return db
}
