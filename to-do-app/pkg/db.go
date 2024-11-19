package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	// Kết nối MySQL
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "to_do_list",
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// case connect error
	if err != nil || db.Ping() != nil {
		fmt.Println("Connected to MySQL ERROR!")

		return nil, err
	}

	// case OK
	DB = db

	return db, nil
}
