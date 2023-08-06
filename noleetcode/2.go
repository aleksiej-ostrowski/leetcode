/*

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const N_STRS = 100000

type MyError struct {
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s", e.What)
}

func checkUrl(st string) error {
	if len(st) < 10 {
		return nil
	}
	return &MyError{
		"ok",
	}
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var counter int64 = 0
	data := make(chan string)

	var wg sync.WaitGroup

	start := time.Now()

	go func() {
		for ind := 0; ind < N_STRS; ind++ {
			data <- RandStringRunes(rand.Intn(20))
		}
		defer close(data)
	}()

	for ind := 0; ind < 100; ind++ {
		wg.Add(1)
		go func() {
			var sm int64 = 0
			for el := range data {
				if err := checkUrl(el); err != nil {
					sm += 1
				}
			}
			atomic.AddInt64(&counter, sm)
			defer wg.Done()
		}()
	}

	wg.Wait()

	duration := time.Since(start)

	fmt.Println("counter = ", counter)
	fmt.Println("time = ", duration)
}
