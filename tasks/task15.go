package main

import (
    "fmt"
)

var justString string

func someFunc() {
    // Если не создать новую копию для среза
    // то GC не сможет удалить большую строку
    // которая неэффективно используется, так как
    // на нее будет указатель из justString:e
    v := createHugeString(1 << 10)
    // Создав новую копию массива для среза
    // мы удалили указатель с v и теперь его
    // можно удалить GC, а мы используем и храним
    // в памяти только нужную часть строки.
    justString = append(justString, v[:100])
}

func main() {
    someFunc()
}
