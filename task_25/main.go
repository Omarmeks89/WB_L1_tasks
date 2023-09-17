package main

import (
    "time"
    "fmt"
)

func sleep(t time.Duration) {
    <-time.After(t)
}

func main() {
    fmt.Printf("Start at %s\n", time.Now())
    sleep(5 * time.Second)
    fmt.Printf("Done at %s\n", time.Now())
}
