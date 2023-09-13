package main

import (
    "fmt"
)

func ReadFromArr(arr *[]int, ch chan<- int) {
    go func(c chan<- int, arr *[]int) {
        defer close(c)
        for _, num := range *arr {
            c<- num
        }
        return
    }(ch, arr)
}

func Square(nums <-chan int, sqrs chan<- int) {
    go func(rd <- chan int, wr chan<- int) {
        defer close(wr)
        for num := range rd {
            wr<- num * num
        }
        return
    }(nums, sqrs)
}

func PrintNums(c <-chan int) {
    for sqr := range c {
        fmt.Printf("Square = %d\n", sqr)
    }
    return
}

func main() {
    nums := make(chan int)
    squares := make(chan int)
    arr := &[]int{2, 4, 6, 8, 10, 12, 14, 16, 18}
    ReadFromArr(arr, nums)
    Square(nums, squares)
    PrintNums(squares)
}
