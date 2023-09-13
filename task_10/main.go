package main

import (
    "fmt"
    "sync"
    "errors"
)

var (
    ItemNotFound = errors.New("Item not found in set")
)

type TemperatureSet struct {
    set map[float64]struct{}
    lock sync.RWMutex
}

func (s *TemperatureSet) Contains(item float64) bool {
    //...
    s.lock.RLock()
    defer s.lock.RUnlock()
    _, contains := s.set[item]
    return contains
}

func (s *TemperatureSet) Set() map[float64]struct{} {
    //... 
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.set
}

func (s *TemperatureSet) Add(item float64) bool {
    //...
    s.lock.Lock()
    defer s.lock.Unlock()
    if _, contains := s.set[item]; contains == true {
        return false
    }
    s.set[item] = struct{}{}
    return true
}

func (s *TemperatureSet) Remove(item float64) (bool, error) {
    //...
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.Contains(item) {
        delete(s.set, item)
        return true, nil
    }
    return false, ItemNotFound
}

func NewTemperatureSet() *TemperatureSet {
    return &TemperatureSet{lock: sync.RWMutex{}, set: make(map[float64]struct{})}
}

type TemprMap struct {
    subsets map[int]*TemperatureSet
    lock sync.RWMutex
}

func (tm *TemprMap) ConvertToKey(temp float64) int {
    tmp := int(temp)
    return tmp - (tmp % 10)
}

func (tm *TemprMap) Add(temp float64) bool {
    //...
    tm.lock.Lock()
    defer tm.lock.Unlock()
    tmp := tm.ConvertToKey(temp)
    if _, contains := tm.subsets[tmp]; contains != true {
        subset := NewTemperatureSet()
        tm.subsets[tmp] = subset
    }
    return tm.subsets[tmp].Add(temp)
}

func (tm *TemprMap) PrintMap() {
    tm.lock.RLock()
    defer tm.lock.RUnlock()
    for idx, tset := range tm.subsets {
        fmt.Printf("TEMP.IDX %d: %v\n", idx, tset.Set())
    }
}

func NewTemperMap() *TemprMap {
    //...
    return &TemprMap{subsets: make(map[int]*TemperatureSet), lock: sync.RWMutex{}}
}

func main() {
    var wg sync.WaitGroup
    tmap := NewTemperMap()
    TVolatile := []float64{-21.0, -25.3, -15.0, 0.0, 10.0, 13.0, 21.0, 32.0, 24.5, 10.0, -21.0, -39.9}
    for i := 0; i < len(TVolatile); i++ {
        wg.Add(1)
        go func(temp float64, tmp *TemprMap, w *sync.WaitGroup) {
            //...
            defer w.Done()
            if added := tmp.Add(temp); added == false {
                fmt.Printf("WRK: value [%.1f] not unique.\n", temp)
                return
            }
            fmt.Printf("WRK: value [%.1f] added.\n", temp)
            return
        }(TVolatile[i], tmap, &wg)
    }
    wg.Wait()
    tmap.PrintMap()
    return
}

