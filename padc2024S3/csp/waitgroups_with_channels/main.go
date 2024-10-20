package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    result := id * 2

    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

    ch <- result

    fmt.Printf("Worker_%d Done\n", id)
}

func main() {
    numWorkers := 3
    ch := make(chan int, numWorkers)
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(i, ch, &wg)
    }

    go func() {
        defer close(ch)
        
        wg.Wait()
        fmt.Println("Closing the channel")
    }()

    for result := range ch {
        fmt.Printf("Main result: %d\n", result)
    }
}
