package main

import (
    "fmt"
    "time"
    "context"
    "math/rand"
)

// функция-писатель, которая будет писать данные в канал
func ProduceData(ctx *context.Context, c chan<- int) {
    for {
        select {
        case <-(*ctx).Done():
            return
        case c<- rand.Intn(1000):
        }
    }
}

// функция, читающая данные из канала
func ReceiveData(ctx *context.Context, c <-chan int) {
    for {
        // Используем select для отслеживания
        // таймаута контекста или появления
        // данных в канале.
        select {
        case <-(*ctx).Done():
            return
        case d := <-c:
            fmt.Printf("Received [%d]\n", d)
        }
    }
}

func main() {
    // В качестве инструмента для отмены был выбран
    // context как универсальный и расширяемый инструмент.
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    chn := make(chan int)
    go ProduceData(&ctx, chn)
    ReceiveData(&ctx, chn)
}
