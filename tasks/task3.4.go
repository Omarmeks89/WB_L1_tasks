package main

import (
    "fmt"
    "sync/atomic"
    "sync"
)

// Решение с использование atomic
// atomic хорошо подходит для
// конкурентной работы со счетчиками
// так как гарантирует консистентность 
// операции для примитивных типов.
type aInt32 struct {
    val int32
}

// Создаем методы для записи и чтоения
// значения и конструктор типа.
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
    // Используем WaitGroup для
    // синхронизации
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
    // Ждем пока все горутины завершат работу
    wg.Wait()
    fmt.Printf("Atomic sum of squares = %d\n", sum.GetValue())
}
