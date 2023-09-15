package main

import (
    "fmt"
    "sync"
    "errors"
)

var (
    TypeError = errors.New("TypeError")
    ItemNotFoundError = errors.New("Item not found")
)

type Set interface {
    Set() map[string]struct{}
}

type StringSet struct {
    set map[string]struct{}
    lock sync.RWMutex
}

func (s *StringSet) Contains(item string) bool {
    //...
    s.lock.RLock()
    defer s.lock.RUnlock()
    _, contains := s.set[item]
    return contains
}

func (s *StringSet) Set() map[string]struct{} {
    //...
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.set
}

func (s *StringSet) Add(item string) (bool, error) {
    //...
    s.lock.Lock()
    defer s.lock.Unlock()
    if _, contains := s.set[item]; contains {
        return false, ItemNotFoundError
    }
    s.set[item] = struct{}{}
    return true, nil
}

func (s *StringSet) FromArray(arr *[]string) *StringSet {
    //...
    newset := NewStringSet()
    for _, item := range *arr {
        newset.Add(item)
    }
    return newset
}

func NewStringSet() *StringSet {
    return &StringSet{lock: sync.RWMutex{}, set: make(map[string]struct{})}
}

func PrintSet(set Set) {
    fmt.Printf("Set: {")
    for item, _ := range set.Set() {
        fmt.Printf("%#v ,", item)
    }
    fmt.Printf("}\n")
}

func main() {
    seq := []string{"cat", "cat", "dog", "cat", "tree"}
    set := NewStringSet().FromArray(&seq)
    PrintSet(set)
}

