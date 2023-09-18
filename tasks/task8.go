package main

import (
    "fmt"
)

func SetBit(num *int64, bitpos int, val int) {
    // Изменяем значение нужного бита, при этом
    // если значение этого бита = нужному, мы
    // не изменяем текущее значение.
    var v int64 = int64(val)
    var p int64 = int64(bitpos)

    bitval := v << p
    // num | bitval => 01101111 | 00010000 => 01111111
    // ^bitval ^ 00010000 => 11101111 ^ 00010000 => 11111111
    // 01111111 & 11111111 => 01111111 (то число, которое нам и нужно)
    *num = (*num | bitval) & (^bitval ^ (1 << p))
}

func main() {
    var num int64 = 64
    var negnum int64 = -128
    fmt.Println("Positive numbers...")
    fmt.Printf("Before = %d\n", num)
    SetBit(&num, 8, 1)
    fmt.Printf("After = %d\n", num)
    fmt.Println("Negative numbers...")
    fmt.Printf("Before = %d\n", negnum)
    SetBit(&negnum, 6, 1)
    fmt.Printf("After = %d\n", negnum)
}
