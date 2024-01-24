package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "159357"
	dbname   = "go"
)

var db *sql.DB

func init() {
	var err error

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}

func main() {
	clearDatabase(db)
	channel := make(chan User, 10_000)
	var wg sync.WaitGroup

	wg.Add(1)
	go generate(10*time.Minute, channel, &wg)

	wg.Add(1)
	go worker2(1*time.Minute, 10, channel, &wg)

	wg.Wait()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func clearDatabase(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM "user"."user"`)
	CheckError(err)
}
