package main

import (
    "fmt"
    "sync"
)

//
// create a type for accumulate
// sum of squares
//
type Summator struct {
    sum int
    lock *sync.RWMutex
}

func (s *Summator) Add(num int) {
    // 
    // we lock obj on read & write
    // for save consistency
    //
    s.lock.Lock()
    defer s.lock.Unlock()
    s.sum += num
}

func (s *Summator) GetSum() int {
    //
    // now we lock writer
    // but we can read.
    //
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.sum
}

func main2() {
    var wg sync.WaitGroup
    summ := &Summator{0, &sync.RWMutex{}}

    arr := [...]int{2, 4, 6, 8, 10}

    for i := 0; i < len(arr); i++ {
        wg.Add(1)
        go func(num int, res *Summator, w *sync.WaitGroup) {
            defer w.Done()
            res.Add(num * num)
        }(arr[i], summ, &wg)
    }
    wg.Wait()
    fmt.Printf("Sum of squares = %d\n", summ.GetSum())
}
