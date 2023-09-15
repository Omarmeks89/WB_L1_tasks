package main

import (
    "sync/atomic"
    "sync"
    "fmt"
)

type Total struct {
    addr int32
}

func (t *Total) Increment() {
    atomic.AddInt32(&t.addr, 1)
}

func (t *Total) Value() int {
    val := atomic.LoadInt32(&t.addr)
	return int(val)
}

func main() {
    //...
    var wg sync.WaitGroup
    var total Total

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                total.Increment()
            }
        }()
    }
    wg.Wait()
    fmt.Printf("Total increments = %d\n", total.Value())
}
