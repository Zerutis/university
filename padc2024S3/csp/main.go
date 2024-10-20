package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))

	defer close(out)

	for _, n := range nums {
		out <- n
	}
	return out
}

func main() {
	in := gen(2, 3, 4, 5, 6)

	for value := range in {
		fmt.Println(value)
	}
}
