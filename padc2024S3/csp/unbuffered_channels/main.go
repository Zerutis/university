package main

import (
	"fmt"
	"runtime"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format(time.TimeOnly)
}

func write(value int, ch chan int) {
	for i := 0; i < value; i++ {
		fmt.Printf("%s: bWRITE: %d\n", formatTime(time.Now()), i)
		
		ch <- i // Writing to channel
		
		time.Sleep(time.Second * 1)

		fmt.Printf("%s: aWRITE: %d\n\n", formatTime(time.Now()), i)
	}
}

func read(id string, ch chan int) {
	for value := range ch { // Reading from channel
		fmt.Printf("%s: |%s| READ: %d\n", formatTime(time.Now()), id, value)
	}
}

func main() {
	ch := make(chan int) // Using "make(chan <type>)" we can create unbuffered channel
	
	fmt.Printf("processes: %d\n", runtime.GOMAXPROCS(0))

	go read("reader0", ch)
	go read("reader1", ch)
	go read("reader2", ch)
	go read("reader3", ch)
	go read("reader4", ch)
	go read("reader5", ch)
	go read("reader6", ch)
	go read("reader7", ch)
	go read("reader8", ch)
	go read("reader9", ch)
	go write(5, ch)
	go write(5, ch)
	
	time.Sleep(time.Second * 10)

	close(ch)
}
