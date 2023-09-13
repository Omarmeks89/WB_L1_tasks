package main

import (
    "fmt"
)

func main1() {
    arr := [...]int{2, 4, 6, 8, 10}
    var sum_of_squares int
    squares := make(chan int)
    defer close(squares)

    for i := 0; i < len(arr); i++ {
        go func(num int, res chan<- int) {
            res<- num * num
        }(arr[i], squares)
    }
    for j := 0; j < len(arr); j++ {
        sum_of_squares += <-squares
    }
    fmt.Printf("Sum of squares = %d\n", sum_of_squares)
}
