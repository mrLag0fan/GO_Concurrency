package main

import (
	"GO_Concurrency/internal"
	"GO_Concurrency/pkg/database"
	"sync"
	"time"
)

func main() {
	database.ClearDatabase(database.DB)
	channel := make(chan internal.User, 10_000)
	var wg sync.WaitGroup

	wg.Add(1)

	go internal.Generate(10*time.Minute, channel, &wg)

	wg.Add(1)
	go internal.Worker2(1*time.Minute, 10, channel, &wg)

	wg.Wait()
}
