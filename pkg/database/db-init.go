package database

import (
	"GO_Concurrency/pkg/errors"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "159357"
	dbname   = "go_user_db"
)

var DB *sql.DB

func init() {
	var err error

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}

func ClearDatabase(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM "go_user"`)
	errors.CheckError(err)
}
