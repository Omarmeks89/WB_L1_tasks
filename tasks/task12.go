package main

import (
    "fmt"
    "errors"
)

var (
    TypeError = errors.New("TypeError")
    ItemNotFoundError = errors.New("Item not found")
)

// создаем тип множества 
// реализуем нужное нам поведения
// для создания и модификации множества
type StringSet struct {
    set map[string]struct{}
}

func (s *StringSet) Set() map[string]struct{} {
    return s.set
}

// если добавляемый элемент не содержится
// вл множестве - добавляем его
func (s *StringSet) Add(item string) (bool, error) {
    if _, contains := s.set[item]; contains {
        return false, ItemNotFoundError
    }
    s.set[item] = struct{}{}
    return true, nil
}

// строим множество из массива значений
func (s *StringSet) FromArray(arr *[]string) *StringSet {
    newset := NewStringSet()
    for _, item := range *arr {
        newset.Add(item)
    }
    return newset
}

func NewStringSet() *StringSet {
    return &StringSet{make(map[string]struct{})}
}

func main() {
    seq := []string{"cat", "cat", "dog", "cat", "tree"}
    // создаем множество из массива
    set := NewStringSet().FromArray(&seq)
    fmt.Printf("Set: {")
    for item, _ := range set.Set() {
        fmt.Printf("%#v ,", item)
    }
    fmt.Printf("}\n")
}

