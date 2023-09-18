package main

import (
    "fmt"
)

func QuickSort(arr *[]int) {
    // O(n log n) по памяти
    // если массив пустой, либо только с одним значением
    // выходим
    var minimal, maximal, result []int
    size := len(*arr)
    if size <= 1 {
	return
    }
    // определяем индекс опорного элемента 
    mid := len(*arr) / 2
    // формируем слайс элементов, которые меньше опорного
    for i := 0; i < size; i++ {
	if (*arr)[i] < (*arr)[mid] {
            minimal = append(minimal, (*arr)[i])
	}
    }
    // формируем слайс элементов, которые больше либо равны
    // опорному кроме него самого
    for j := 0; j < size; j++ {
	if (*arr)[j] >= (*arr)[mid] && j != mid {
            maximal = append(maximal, (*arr)[j])
	}
    }
    // рекурсивно сортируем созданные слайсы
    QuickSort(&minimal)
    QuickSort(&maximal)
    // Копируем результаты в новый слайс
    result = append(result, minimal...)
    result = append(result, (*arr)[mid])
    result = append(result, maximal...)
    // изменяем значение указателя на новое
    *arr = result
}

func main() {
    var a []int = []int{1, 0, 0, 6, 2, -1, -100, -500, 1, 1, 5, 0, -8}
	fmt.Printf("Income: %+v\n", a)
	QuickSort(&a)
	fmt.Printf("Sorted >>> %v\n", a)
}
