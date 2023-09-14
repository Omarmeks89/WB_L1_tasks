package main

import (
    "fmt"
    "math/rand"
    "time"
    "context"
)

func main() {
    //...
    rand.Seed(time.Now().UTC().UnixNano())
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    data := make(chan int)

    go func(ctx *context.Context, c chan<- int) {
        defer close(c)
        for {
            select {
            case c<- rand.Intn(100):
                time.Sleep(1 * time.Second)
            case <-(*ctx).Done():
                return
            }
        }
    }(&ctx, data)

    for n := range data {
        if n == 95 {
            fmt.Printf("Number = %d, let`s done!\n", n)
            return
        }
        fmt.Printf("Number = %d, continue!\n", n)
    }
    fmt.Println("Done by context timeout...")
}


