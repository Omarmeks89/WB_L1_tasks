package main

import (
    "fmt"
    "errors"
)

var (
    IndexError = errors.New("index out of range")
)

// delete i-th element from slice
func Delete(src []int, idx int) ([]int, error) {
    if idx > len(src) - 1 || idx < 0 {
        return []int{}, IndexError
    }
    return append(src[:idx], src[idx + 1:]...), nil
}

func main() {
    slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 13, 14}
    nc, _ := Delete(slc, 0)
    fmt.Println(nc) // [2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 13, 14]
    nc, _ = Delete(nc, 4)
    fmt.Println(nc) // [2, 3, 4, 5, 7, 8, 9, 0, 12, 13, 14]
    nc, _ = Delete(nc, 3)
    fmt.Println(nc) // [2, 3, 4, 5, 7, 9, 0, 12, 13, 14]
    nc, _ = Delete(nc, 5)
    fmt.Println(nc) // [2, 3, 4, 5, 7, 9, 0, 12, 14]
}
