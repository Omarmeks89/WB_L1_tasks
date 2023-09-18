package main

import (
    "fmt"
    "errors"
)

var (
    IndexError = errors.New("index out of range")
)

// проверяем, находится ли индекс в диапазоне,
// и если он в диапазоне, то создаем новый слайс
// копируя все элементы кроме искомого, возвращаем 
// новый слас каждый раз
func Delete(src []int, idx int) ([]int, error) {
    if idx > len(src) - 1 || idx < 0 {
        return []int{}, IndexError
    }
    // в качестве базы используем слайс размером
    // до нужного элемента, копируем туда оставшиеся элементы
    // и поскольку они превысят размер слайса, будет создан
    // новый слайс (и массив под ним)
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
