package main

import (
    "fmt"
    "sync"
)

// Решение с использованием mutex
// без создания типа
func main3() {
    var wg sync.WaitGroup
    // Можем использовать обычный мьютекс
    // так как не будем пытаться одновременно
    // читать и изменять значение
    var mt *sync.Mutex = &sync.Mutex{}
    var sum int

    arr := [...]int{2, 4, 6, 8, 10}

    for i := 0; i < len(arr); i++ {
        wg.Add(1)
        go func(num int, res *int, m *sync.Mutex, w *sync.WaitGroup) {
            m.Lock()
            // Отложенные вызовы будут разрешены
            // как LIFO:
            // 1) Unlock
            // 2) Done
            defer w.Done()
            defer m.Unlock()
            *res += num * num
        }(arr[i], &sum, mt, &wg)
    }
    // Ждем пока все горутины завершат работу
    wg.Wait()
    fmt.Printf("Sum of squares = %d\n", sum)
}
