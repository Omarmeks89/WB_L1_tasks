package main

import (
    "fmt"
    "sync"
    "errors"
)

var (
    ItemNotFound = errors.New("Item not found in set")
)

// Создаем тип множества
// Используем RWMutex так как нам
// нужно конкурентно писать и читать данные
type TemperatureSet struct {
    set map[float64]struct{}
    lock sync.RWMutex
}

func (s *TemperatureSet) Contains(item float64) bool {
    // метод подтверждающий наличие
    // значения во множестве
    // Блокируем писателей, но не читателей
    s.lock.RLock()
    defer s.lock.RUnlock()
    _, contains := s.set[item]
    return contains
}

func (s *TemperatureSet) Set() map[float64]struct{} {
    // Пишем в map
    // Блокируем всех
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.set
}

func (s *TemperatureSet) Add(item float64) bool {
    // Пишем
    // Блокируем всех
    s.lock.Lock()
    defer s.lock.Unlock()
    // не используем Contains так как
    // не сможем взять RLock -> писатель заблокировал
    // будет дедлок
    if _, contains := s.set[item]; contains == true {
        return false
    }
    s.set[item] = struct{}{}
    return true
}

func (s *TemperatureSet) Remove(item float64) (bool, error) {
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

// описываем тип для map
type TemprMap struct {
    subsets map[int]*TemperatureSet
    lock sync.RWMutex
}

// для сортировки чисел по нужным ключам (диапазонам)
// 25.9 - (25.9 % 10) = 20, згачит пишем в 20
func (tm *TemprMap) ConvertToKey(temp float64) int {
    tmp := int(temp)
    return tmp - (tmp % 10)
}

func (tm *TemprMap) Add(temp float64) bool {
    // пишем, все (писатели / читатели) ждут
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
    // блокируем писателей, но можем читать
    // и выводить результаты
    tm.lock.RLock()
    defer tm.lock.RUnlock()
    for idx, tset := range tm.subsets {
        fmt.Printf("TEMP.IDX %d: %v\n", idx, tset.Set())
    }
}

func NewTemperMap() *TemprMap {
    return &TemprMap{subsets: make(map[int]*TemperatureSet), lock: sync.RWMutex{}}
}

func main() {
    // будем писать конкурентно
    // Синхронизируем по WaitGroup
    var wg sync.WaitGroup
    tmap := NewTemperMap()
    TVolatile := []float64{-21.0, -25.3, -15.0, 0.0, 10.0, 13.0, 21.0, 32.0, 24.5, 10.0, -21.0, -39.9}
    for i := 0; i < len(TVolatile); i++ {
        // workers = len(numbers)
        wg.Add(1)
        go func(temp float64, tmp *TemprMap, w *sync.WaitGroup) {
            defer w.Done()
            // пишем в  map
            if added := tmp.Add(temp); added == false {
                fmt.Printf("WRK: value [%.1f] not unique.\n", temp)
                return
            }
            fmt.Printf("WRK: value [%.1f] added.\n", temp)
            return
        }(TVolatile[i], tmap, &wg)
    }
    // Ждем пока все горутины отработают
    wg.Wait()
    tmap.PrintMap()
    return
}

