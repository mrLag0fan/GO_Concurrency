package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(period time.Duration, numGoroutines int, channel chan User, wg *sync.WaitGroup) {
	defer wg.Done()
	uptimeTicker := time.NewTicker(period)
	var dataSlice []User
	for {
		select {
		case data, ok := <-channel:
			if !ok {
				fmt.Println("Finishing..........")
				post(dataSlice, numGoroutines)
				fmt.Println("Done.")
				return
			}
			dataSlice = append(dataSlice, data)
		case <-uptimeTicker.C:

			fmt.Println("Posting..........")
			post(dataSlice, numGoroutines)
			dataSlice = nil
			fmt.Println("Done..")
		}
	}
}

func post(dataSlice []User, numGoroutines int) {
	recordsPerGoroutine := len(dataSlice) / numGoroutines
	var goroutineWG sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		goroutineWG.Add(1)

		go func(startIndex int64) {
			defer goroutineWG.Done()
			tx, err := db.Begin()
			endIndex := startIndex + int64(recordsPerGoroutine)
			for j := startIndex; j < endIndex; j++ {
				_, err = tx.Exec(`INSERT INTO "user"."user" VALUES ($1, $2)`,
					dataSlice[j].Name,
					dataSlice[j].Comment)
				CheckError(err)
			}
			err = tx.Commit()
			CheckError(err)
		}(int64(i * recordsPerGoroutine))

	}
	goroutineWG.Wait()
}
