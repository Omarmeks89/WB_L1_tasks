package main

import (
    "sync"
    "fmt"
)

type Counter struct {
    lock sync.RWMutex
    stat map[string]int
}

func (c *Counter) Increment(str string) {
    defer c.lock.Unlock()
    c.lock.Lock()
    c.stat[str]++
}

func (c *Counter) Value(str string) int {
    defer c.lock.RUnlock()
    c.lock.RLock()
    return c.stat[str]
}

func (c *Counter) Range(fn func(key string, val int)) {
    defer c.lock.RUnlock()
    c.lock.RLock()
    for key, val := range c.stat {
    	fn(key, val)
    }
}

func NewCounter() *Counter {
    return &Counter{sync.RWMutex{}, make(map[string]int)}
}

func main() {
    //...
    var wg sync.WaitGroup
    var arr []string = []string{"A", "B", "C", "D"}
    c := NewCounter()

    increment := func(item string, times int) {
        defer wg.Done()
        for ; times > 0; times-- {
            c.Increment(item)
        }
    }

    for i := 0; i < len(arr); i++ {
        wg.Add(1)
        go increment(arr[i], 100)
    }

    wg.Wait()
    fmt.Print("{ ")
    c.Range(func(key string, val int) {
        fmt.Printf("%s:%d ", key, val)
    })
    fmt.Println("}")
}

