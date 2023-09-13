package main

import (
    "fmt"
    "time"
    "context"
    "math/rand"
)

func ProduceData(ctx *context.Context, c chan<- int) {
    //...
    for {
        select {
        case <-(*ctx).Done():
            return
        case c<- rand.Intn(1000):
        }
    }
}

func ReceiveData(ctx *context.Context, c <-chan int) {
    //...
    for d := range c {
        select {
        case <-(*ctx).Done():
            return
        default:
            fmt.Printf("Received [%d]\n", d)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    chn := make(chan int)
    go ProduceData(&ctx, chn)
    ReceiveData(&ctx, chn)
}
