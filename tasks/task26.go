package main

import (
    "fmt"
    "os"
    "unicode"
    "unicode/utf8"
    "errors"
)

var (
    InvalidLiteral = errors.New("Numeric symbols not allowed")
    RepeatedLiteral = errors.New("Repeated literals found")
)

// определяем уникальные символы строке
// и выходим если попадаются повторы
// не чувствительна к регистру.
// При обнаружении чисел в строке возвращает
// ошибку, при обнаружении повторов в строке
// также возвращает ошибку.
func IsUnique(str string) (bool, error) {
    unique := make(map[rune]struct{}, utf8.RuneCountInString(str))
    for _, r := range []rune(str) {
        if unicode.IsDigit(r) {
            return false, InvalidLiteral
        }
        symb := unicode.ToLower(r)
        if _, contains := unique[symb]; !contains {
            unique[symb] = struct{}{}
        } else {
            return false, RepeatedLiteral
        }
    }
    return true, nil
}

func main() {
    str := os.Args[1]
    unique, err := IsUnique(str)
    if err != nil {
        fmt.Printf("%s - %v. ERR: %s\n", str, unique, err.Error())
        os.Exit(1)
    }
    fmt.Printf("%s - %v\n", str, unique)
}
