package main

import (
    "fmt"
)

// для бинарного поиска массив должен быть
// предварительно отсортирован
func BinarySearch(num int, arr []int) bool {
    // Сложность O(n log n)
    // по памяти O(1)
    // определяем индекс опорного элемента
    mid := len(arr) / 2
    // если осталось одно значение и оно
    // не соответствует искомому, вернем false
    if num != arr[mid] && len(arr) <= 1 {
        return false
    }
    // если значение искомого элемента
    // меньше опорного, ищем в меньшей (по значениям)
    // части массива
    if num < arr[mid] {
        return BinarySearch(num, arr[:mid])
    }
    // Если элемент больше - то в большей части
    if num > arr[mid] {
        return BinarySearch(num, arr[mid:])
    }
    // Если ни одно условие не сработало - мы нашли
    // нужный элемент.
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
