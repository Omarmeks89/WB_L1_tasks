package main

import (
    "fmt"
    "unicode/utf8"
)

func ReverseStr(str *string) {
    strlen := utf8.RuneCountInString(*str)
    // O(n)
    dest := []rune(*str)
    i, j := 0, strlen - 1
    // O(n)
    for ; i < j; {
        dest[i], dest[j] = dest[j], dest[i]
        i++
        j--
    }
    // O(n)
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
