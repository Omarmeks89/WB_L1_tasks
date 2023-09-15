package main

import (
    "fmt"
    "errors"
)

var (
    IndexError = erorrs.New("Index out of range.")
)

// Was deleted global variable 
// separate hugestring creation from
// slicing. Use hugestring as a ptr.
func createHugeString(size int) *string {
    var hugeSrc string
    return &hugeSrc
}

// magic-numbers was replaced on named identifyers
// now we can directly set wished slice size
func someFunc(source *string, lim int) (string, error) {
    if lim > len(*source)-1 {
        return "", IndexError
    }
    return (*source)[:lim], nil
}

func main() {
    var justString string
    hs := createHugeString()
    // now we can split from hugestring
    // without creation it any time
    justString, err := someFunc(hs, 5)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(justString)
}

