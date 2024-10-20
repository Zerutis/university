package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() { // send
		ch <- "Hello world!"
	}()
	
	fmt.Println(<-ch) // receive
}
