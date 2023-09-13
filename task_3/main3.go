package main

import (
    "fmt"
    "sync"
)

func main3() {
    var wg sync.WaitGroup
    var mt *sync.Mutex = &sync.Mutex{}
    var sum int

    arr := [...]int{2, 4, 6, 8, 10}

    for i := 0; i < len(arr); i++ {
        wg.Add(1)
        go func(num int, res *int, m *sync.Mutex, w *sync.WaitGroup) {
            m.Lock()
            defer w.Done()
            defer m.Unlock()
            *res += num * num
        }(arr[i], &sum, mt, &wg)
    }
    wg.Wait()
    //
    // we didn`t use RWMutex here bcs we just finished here
    //
    fmt.Printf("Sum of squares = %d\n", sum)
}
