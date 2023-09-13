package main

import (
    "fmt"
)

func SetBit(num *int64, bitpos int, val int) {
    //...
    var v int64 = int64(val)
    var p int64 = int64(bitpos)

    bitval := v << p
    *num = (*num | bitval) & (^bitval ^ ((v | 1) << p))
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
