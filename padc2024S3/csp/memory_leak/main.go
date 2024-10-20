package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int, 1)

	go func() { // Single sender
		// defer close(ch) // DON'T forget to close the channel

		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	go func() { // Multiple receivers
		for i := 0; i < 3; i++ {
			go func(id int) {
				for v := range ch {
					fmt.Printf("goroutine_%d: %d\n", id, v)
				}
			}(i)
		}
	}()

	time.Sleep(time.Second * 1)

	fmt.Printf("Active goroutines: %d", runtime.NumGoroutine())
}
