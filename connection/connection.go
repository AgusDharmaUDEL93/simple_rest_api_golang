package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectAPI() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajarmysql")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, err
}
