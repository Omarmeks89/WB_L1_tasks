package main

import (
    "fmt"
)

func DetermineType(t interface{}) {
    // используем type assertion для
    // получения типа параметра.
    // Далее с помощью switch
    // находим нужную ветку для работы с типом.
    // Типов каналов можем быть множество, поэтому
    // в контексте формулироваки текущей задачи
    // они вынесены в блок default
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
