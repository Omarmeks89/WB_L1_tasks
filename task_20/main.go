package main

import (
    "fmt"
    "strings"
)

const (
    BySpace string = " "
)

func ReverseWords(str *string) {
    words := strings.Fields(*str)
    i, j := 0, len(words) - 1
    for ; i < j; {
        words[i], words[j] = words[j], words[i]
        i++
        j--
    }
    *str = strings.Join(words, BySpace)
}

func main() {
    phrases := []string{
        "one two three",
        "go on",
        "oneline",
        "что-то по русски",
    }
    for _, phr := range phrases {
        orig := phr
        ReverseWords(&phr)
        fmt.Printf("%s -> %s\n", orig, phr)
    }
}
