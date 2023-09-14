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
    timeout := 5 * time.Second

    go func(c chan<- int) {
        defer close(c)
        for {
            tm := rand.Intn(15)
            time.Sleep(time.Duration(int64(tm)) * time.Second)
            c<- tm
            fmt.Println("WRK: sent...")
        }
    }(data)

    for {
        // 
        // in this case we can`t use <range>
        // bcs we`re blocking on waiting data from chnl
        // and timeout can done tahis time
        // and after that we can`t repeat this timeout
        //
        fmt.Printf("TIME: %v\n", time.Now().UTC())
        select {
        case num := <-data:
            fmt.Printf("NUM = %d. CONTINUE\n", num)
        case <-time.After(timeout):
            fmt.Println("TIMEOUT...")
            return
        }
    }
}
