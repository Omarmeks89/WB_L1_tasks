package main

import (
    "fmt"
)

func main() {
    a := 10
    b := 5
    fmt.Printf("a = %d, b = %d\n", a, b)
    a, b = b, a
    fmt.Printf("a = %d, b = %d\n", a, b)
}
