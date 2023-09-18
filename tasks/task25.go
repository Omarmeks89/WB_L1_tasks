package main

import (
    "time"
    "fmt"
)

// для реализации воспользуемся
// функцией After(), которая принимает 
// нужный интервал времени, и возвращает канал
// в который она по истечении времени, запишет 
// данные, после чего наша функция сможет завершиться
func sleep(t time.Duration) {
    <-time.After(t)
}

func main() {
    fmt.Printf("Start at %s\n", time.Now())
    sleep(5 * time.Second)
    fmt.Printf("Done at %s\n", time.Now())
}
