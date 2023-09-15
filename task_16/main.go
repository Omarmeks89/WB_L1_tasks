package main

import (
    "fmt"
)

func QuickSort(arr *[]int) {
    var minimal, maximal, result []int
    size := len(*arr)
    if size <= 1 {
	return
    }
    mid := len(*arr) / 2
    for i := 0; i < size; i++ {
	if (*arr)[i] < (*arr)[mid] {
            minimal = append(minimal, (*arr)[i])
	}
    }
    for j := 0; j < size; j++ {
	if (*arr)[j] >= (*arr)[mid] && j != mid {
            maximal = append(maximal, (*arr)[j])
	}
    }
    QuickSort(&minimal)
    QuickSort(&maximal)
    result = append(result, minimal...)
    result = append(result, (*arr)[mid])
    result = append(result, maximal...)
    *arr = result
}

func main() {
    //...
    var a []int = []int{1, 0, 0, 6, 2, -1, -100, -500, 1, 1, 5, 0, -8}
	fmt.Printf("Income: %+v\n", a)
	QuickSort(&a)
	fmt.Printf("Sorted >>> %v\n", a)
}
