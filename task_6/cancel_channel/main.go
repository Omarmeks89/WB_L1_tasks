package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //...
    rand.Seed(time.Now().UTC().UnixNano())
    data := make(chan int)
    cancel := make(chan struct{})
    //
    // we can use defer:
    // >>> defer close(cancel)
    // result will be the same as direct closing
    // in main loop
    //

    go func(c chan<- int, canc <-chan struct{}) {
        for {
            select {
            case c<- rand.Intn(20):
            case <-canc:
                return
            }
        }
    }(data, cancel)

    for num := range data {
        if num == 10 {
            fmt.Printf("NUM = %d. EXIT\n", num)
            close(cancel)
            return
        }
        fmt.Printf("NUM = %d. CONTINUE\n", num)
    }
}
