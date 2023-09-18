package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Использование таймаута для отмены горутины.
    // Может быть полезно если нужно отслеживать время
    // работы конкретной/ых горутин и использование
    // context кажется избыточным.
    rand.Seed(time.Now().UTC().UnixNano())
    data := make(chan int)
    timeout := 5 * time.Second

    go func(c chan<- int) {
        defer close(c)
        tm := rand.Intn(15)
        time.Sleep(time.Duration(int64(tm)) * time.Second)
        fmt.Println("WRK: work done")
    }(data)

    fmt.Printf("TIME: %v\n", time.Now().UTC())
    select {
    case <-data:
        fmt.Printf("worker done, ok")
    case <-time.After(timeout):
        fmt.Println("TIMEOUT...")
        return
    }
}
