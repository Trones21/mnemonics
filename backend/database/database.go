package database

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

// Test gitignore
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mnemonicsdb")
	if err != nil {
		log.Fatal(err)
	}
}
