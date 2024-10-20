package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func main() {
	for i:= 0 ; i < 5; i++ {
		go greet(fmt.Sprintf("goroutine_%d", i)) // Using "go" key, we can start a lightweight thread called goroutine 
	}

	time.Sleep(time.Millisecond * 500)

	greet("main")
}
