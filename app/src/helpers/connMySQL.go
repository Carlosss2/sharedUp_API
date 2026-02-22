package helpers

import (
	"database/sql"
	"fmt"
	"os"
	"time"
_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)


func ConectToMySQL() (db *sql.DB, err error){
	loadVerify := godotenv.Load()

	if loadVerify != nil {
		return nil, loadVerify
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	//Manejo de conexión poll
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(16)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	print("Conexión establecida con éxito")
	return db, nil

}