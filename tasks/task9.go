package main

import (
    "fmt"
)

// Читаем числа из массива и пишем в канал
// Когда числа кончатся - закрываем канал, сигнали
// зируя что данных больше не будет
func ReadFromArr(arr *[]int, ch chan<- int) {
    go func(c chan<- int, arr *[]int) {
        defer close(c)
        for _, num := range *arr {
            c<- num
        }
        return
    }(ch, arr)
}

// Читаем числа из канала чисел, возводим из в квадрат
// и записываем в канал квадратов чисел.
// Читаем канал с помощью range - так мы поймем что канал закрыт
// и можно завершать работу.
// Канал квадратов (так как мы писатель) закрываем.
func Square(nums <-chan int, sqrs chan<- int) {
    go func(rd <- chan int, wr chan<- int) {
        defer close(wr)
        for num := range rd {
            wr<- num * num
        }
        return
    }(nums, sqrs)
}

// Синхронная функция-читатель
// Читаем квадраты чисел и выводим в sdtout
// Используем range, чтобы понять что канал закрыт
// и можно завершать рвботу
func PrintNums(c <-chan int) {
    for sqr := range c {
        fmt.Printf("Square = %d\n", sqr)
    }
    return
}

func main() {
    // создаем канал для чисел и для
    // квадратов этих чисел
    nums := make(chan int)
    squares := make(chan int)
    arr := &[]int{2, 4, 6, 8, 10, 12, 14, 16, 18}
    ReadFromArr(arr, nums)
    Square(nums, squares)
    PrintNums(squares)
}
