package main

import (
    "fmt"
    "sync"
    "errors"
)

var (
    ItemNotFound = errors.New("Item not found.")
)

type IntSet struct {
    set map[int]struct{}
    lock sync.RWMutex
}

func (s *IntSet) Contains(item int) bool {
    //...
    s.lock.RLock()
    defer s.lock.RUnlock()
    _, contains := s.set[item]
    return contains
}

func (s *IntSet) Set() map[int]struct{} {
    //...
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.set
}

func (s *IntSet) Add(item int) bool {
    //...
    s.lock.Lock()
    defer s.lock.Unlock()
    if _, contains := s.set[item]; contains == true {
        return false
    }
    s.set[item] = struct{}{}
    return true
}

func (s *IntSet) Intersect(ns *IntSet) *IntSet {
    //...
    s.lock.RLock()
    defer s.lock.RUnlock()
    intersection := NewIntSet()
    for item, _ := range ns.Set() {
        if s.Contains(item) {
            intersection.Add(item)
        }
    }
    return intersection
}

func (s *IntSet) FromArray(arr *[]int) *IntSet {
    //...
    newset := NewIntSet()
    for _, item := range *arr {
        newset.Add(item)
    }
    return newset
}

func (s *IntSet) Remove(item int) (bool, error) {
    //...
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.Contains(item) {
        delete(s.set, item)
        return true, nil
    }
    return false, ItemNotFound
}

func NewIntSet() *IntSet {
    return &IntSet{lock: sync.RWMutex{}, set: make(map[int]struct{})}
}

func main() {
    //...
    arr1 := []int{2, 3, 12, 56, 0, -1, 5, 67, 98}
    arr2 := []int{3, 56, 4, 34, 2, -1, 0, -13, 98, 100}
    set1 := NewIntSet().FromArray(&arr1)
    set2 := NewIntSet().FromArray(&arr2)
    intersect := set1.Intersect(set2)
    fmt.Printf("Set1: %v\n", set1.Set())
    fmt.Printf("Set2: %v\n", set2.Set())
    fmt.Printf("Set1 & Set2: %v\n", intersect.Set())
}
