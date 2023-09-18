package main

import (
    "sync/atomic"
    "sync"
    "fmt"
)

// создаем тип для счетчика
type Total struct {
    addr int32
}

// atomic не требует блокировок для своей
// работы. Изменение значения осуществляется
// за одну инструкцию процессора.
// работает с примитивными типами данных
func (t *Total) Increment() {
    atomic.AddInt32(&t.addr, 1)
}

func (t *Total) Value() int {
    val := atomic.LoadInt32(&t.addr)
	return int(val)
}

func main() {
    // синхронизируем через WaitGroup
    var wg sync.WaitGroup
    var total Total

    for i := 0; i < 10; i++ {
        // пишем в 10 горутин
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                total.Increment()
            }
        }()
    }
    // ждем пока все горутины завершат работу
    wg.Wait()
    fmt.Printf("Total increments = %d\n", total.Value())
}
