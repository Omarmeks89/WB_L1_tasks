package main

import (
    "fmt"
    "sync/atomic"
    "sync"
)

//
type aInt32 struct {
    val int32
}

func (a *aInt32) Add(num int) {
    atomic.AddInt32(&a.val, int32(num))
}

func (a *aInt32) GetValue() int {
    return int(atomic.LoadInt32(&a.val))
}

func New_aInt32(value int) *aInt32 {
    return &aInt32{int32(value)}
}

func main4() {
    var wg sync.WaitGroup
    var sum *aInt32 = New_aInt32(0)
    
    arr := [...]int{2, 4, 6, 8, 10}

    for i := 0; i < len(arr); i++ {
        wg.Add(1)
        go func(num int, s *aInt32, w *sync.WaitGroup) {
            defer w.Done()
            s.Add(num * num)
        }(arr[i], sum, &wg)
    }
    wg.Wait()
    fmt.Printf("Atomic sum of squares = %d\n", sum.GetValue())
}
