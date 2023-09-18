package main

import (
    "fmt"
    "math/rand"
    "time"
    "context"
)

// Отмена выполнения горутины по контексту.
// Контекст - универсальный инструмент, созданный
// специально для отмены горутин, поодерживает
// отмену по таймауту, дедлайну и напрямую фунцкией отмены.
// Контекст неизменяем + к основному контексту
// при необходимости можно создать новый контекст
// (как правило с более жесткими ограничениями либо
// просто иным способом завершения).
func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    data := make(chan int)

    go func(ctx *context.Context, c chan<- int) {
        // Планируем закрытие канала
        // после завершения горутины
        defer close(c)
        for {
            select {
            case c<- rand.Intn(100):
                time.Sleep(1 * time.Second)
            case <-(*ctx).Done():
                return
            }
        }
    }(&ctx, data)

    // Для обработки ситуации закрытия канала
    // читаем его с помощью range - он выйдет из
    // цикла когда увидит закрытое значение канала
    for n := range data {
        if n == 95 {
            fmt.Printf("Number = %d, let`s done!\n", n)
            return
        }
        fmt.Printf("Number = %d, continue!\n", n)
    }
    fmt.Println("Done by context timeout...")
}


