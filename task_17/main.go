package main

import (
    "fmt"
)

func BinarySearch(num int, arr []int) bool {
    mid := len(arr) / 2
    if num != arr[mid] && len(arr) <= 1 {
        return false
    }
    if num < arr[mid] {
        return BinarySearch(num, arr[:mid])
    }
    if num > arr[mid] {
        return BinarySearch(num, arr[mid:])
    }
    return true
}

func main() {
    var arr []int = []int{-3, 0, 1, 4, 5, 8, 12, 15}
    if found := BinarySearch(0, arr); found == true {
        fmt.Printf("%d found.\n", int(0))
	return
    }
    fmt.Printf("%d not found.\n", int(0))
}
