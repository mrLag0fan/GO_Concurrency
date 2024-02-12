package internal

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name    string
	Comment string
}

func (user User) String() string {
	return fmt.Sprintf("Name: %s \nComment: %s", user.Name, user.Comment)
}

var count = 0

func Generate(duration time.Duration, channel chan User, wg *sync.WaitGroup) {
	for timeout := time.After(duration); ; {
		select {
		case channel <- User{
			Name:    fmt.Sprintf("%s%d", "name", count),
			Comment: fmt.Sprintf("%s%d", "comment", count),
		}:
			count++
			time.Sleep(10)
		case <-timeout:
			close(channel)
			wg.Done()
			fmt.Println("Stop generating")
			return
		}
	}
}
