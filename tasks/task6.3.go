package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Использование канала отмены
// Создается канал в который никто не пишет
// т.е. попытка чтения из него блокирующая и Select
// его пропустит (пойдет к другим каналам)
// При закрытии cancel канала, канал разблокируется
// (читать из закрытого канала можно - это будет пустые и false
// значения) и в этой ветке в горутине стоит оператор return
func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    cancel := make(chan struct{})
    data := make(chan int)

    go func(c chan<- int, canc <-chan struct{}) {
        for {
            select {
            // при определенном условии
            // читатель перестанет читать данные (те
            // эта ветвь заблокируется) и закроет cancel
            // канал (те ветвь выхода разблокируется)
            case c<- rand.Intn(20):
            case <-canc:
                return
            }
        }
    }(data, cancel)

    for num := range data {
        if num == 10 {
            fmt.Printf("NUM = %d. EXIT\n", num)
            close(cancel)
            return
        }
        fmt.Printf("NUM = %d. CONTINUE\n", num)
    }
}
