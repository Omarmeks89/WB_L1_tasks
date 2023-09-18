package main

import (
    "fmt"
)

func partition(arr []int, low, high int) ([]int, int) {
    // находим опорный элемент
    pivot := arr[high]
    i := low
    for j := low; j < high; j++ {
        // если значение меньше выбранного
        // заменяем его на другое
        if arr[j] < pivot {
            arr[i], arr[j] = arr[j], arr[i]
            i++
	}
    }
    // меняем начало и конец местами чтобы
    // позже определить значение по a[j]
    arr[i], arr[high] = arr[high], arr[i]
    // возвращаем слайс и текущую позицию
    // которую впоследствии примем за верхнюю границу
    return arr, i
}

func quickSort(arr []int, low, high int) []int {
    if low < high {
        var p int
        // в цикле делим массив на части для 
        // сортировки и определяем границы частей
        arr, p = partition(arr, low, high)
        arr = quickSort(arr, low, p-1)
        arr = quickSort(arr, p+1, high)
    }
    return arr
}

// пуск сортировки
func quickSortStart(arr []int) []int {
    return quickSort(arr, 0, len(arr)-1)
}

func main() {
    fmt.Println(quickSortStart([]int{5, 1, 20, -1, 0, -12, 104}))
}

