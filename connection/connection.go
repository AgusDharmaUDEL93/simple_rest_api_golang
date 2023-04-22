package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectAPI() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	DB_USERNAME := os.Getenv("DB_USERNAME");	
	DB_PASSWORD := os.Getenv("DB_PASSWORD");	
	DB_HOST := os.Getenv("DB_HOST");	
	DB := os.Getenv("DB");

	db, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+")/"+DB)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, err
}
