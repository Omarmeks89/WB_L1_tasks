package main

import (
    "fmt"
)

func DetermineType(t interface{}) {
    switch typedef := t.(type) {
    case string:
        fmt.Printf("It`s a %T type\n", typedef)
    case int:
        fmt.Printf("It`s a %T type\n", typedef)
    case bool:
        fmt.Printf("It`s a %T type\n", typedef)
    default:
        fmt.Printf("Channel type: %T\n", typedef)
    }
}

func main() {
    //...
    types := map[int]interface{}{
        0:              int(1),
        2:              string("b"),
        3:              make(chan float32),
        4:              1 == 1,
    }
    for _, tp := range types {
        DetermineType(tp)
    }
}
