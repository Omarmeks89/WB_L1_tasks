package main

import (
    "fmt"
    "unicode/utf8"
)

func ReverseStr(str *string) {
    // получаем длину строки через RuneCount
    // так как мы ожидаем символю unicode
    // и функция len предоствит нам информацию о кол-ве
    // байт в строке, а unicode можем быть 2 и 3 байта
    strlen := utf8.RuneCountInString(*str)
    // O(n)
    // создаем массив рун для новой строки (в байт не поместятся)
    // rune = int32
    dest := []rune(*str)
    i, j := 0, strlen - 1
    // O(n)
    for ; i < j; {
        // разворачиваем строку
        dest[i], dest[j] = dest[j], dest[i]
        i++
        j--
    }
    // O(n)
    // переставляем указатель на новую строку
    *str = string(dest)
}

func main() {
    // ...
    strgs := [...]string{
        "first",
        "second",
        "последняя",
    }
    for _, str := range strgs {
        origin := str
        ReverseStr(&str)
        fmt.Printf("%s -> %s\n", origin, str)
    }
}
