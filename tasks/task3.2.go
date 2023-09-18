package main

import (
    "fmt"
    "sync"
)

// Создаем тип-счетчик с примитивом синхронизации
type Summator struct {
    sum int
    lock *sync.RWMutex
}

func (s *Summator) Add(num int) {
    // Lock блокирует
    // возможность чтения и записи для всех,
    // кроме взявшей его горутины.
    s.lock.Lock()
    defer s.lock.Unlock()
    s.sum += num
}

func (s *Summator) GetSum() int {
    // Получать значения можно конкурентно
    // благодаря методу RLock(), который блокирует
    // писателей, но позволяет читать состояние.
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.sum
}

func main2() {
    // Для синхронизации используем WaitGroup
    var wg sync.WaitGroup
    summ := &Summator{0, &sync.RWMutex{}}

    arr := [...]int{2, 4, 6, 8, 10}

    for i := 0; i < len(arr); i++ {
        // увеличиваем счетчик
        wg.Add(1)
        go func(num int, res *Summator, w *sync.WaitGroup) {
            // Планируем декремент счетчика
            // после завершения горутины
            defer w.Done()
            res.Add(num * num)
        }(arr[i], summ, &wg)
    }
    // Ждем пока счетчик не будет равен 0, значит
    // все горутины завершили работу.
    wg.Wait()
    // Смотрим что получилось
    fmt.Printf("Sum of squares = %d\n", summ.GetSum())
}
