package main

import (
    "fmt"
    "sync"
)

type ConcurrentMap struct {
    lock sync.RWMutex
    coll map[string][]string
}

func (c *ConcurrentMap) Add(key, value string) {
    defer c.lock.Unlock()
    c.lock.Lock()
    if _, contains := c.coll[key]; contains == true {
        c.coll[key] = append(c.coll[key], value)
    } else {
        c.coll[key] = []string{value}
    }
    return
}

func (c *ConcurrentMap) Get(key string) []string {
    defer c.lock.RUnlock()
    c.lock.RLock()
    return c.coll[key]
}

func (c *ConcurrentMap) Range(fn func(key string, val []string)) {
    defer c.lock.RUnlock()
    c.lock.RLock()
    for key, val := range c.coll {
	fn(key, val)
    }
}

func NewConcMap() *ConcurrentMap {
    return &ConcurrentMap{sync.RWMutex{}, make(map[string][]string)}
}

func main() {
    var wg sync.WaitGroup
    var mapkeys []string = []string{"A", "B", "C", "D"}
    cmap := NewConcMap()
    wcount := 8
    keys := make(chan string, 4)
    for i := 0; i < 4; i++ {
        keys<- mapkeys[i]
    }
    for j := 0; j < wcount; j++ {
        wg.Add(1)
        go func(w *sync.WaitGroup, k chan string, m *ConcurrentMap, n int) {
            defer w.Done()
            var operations int = 0
            for i := 0; i < 10; i++ {
                key := <-k
                msg := fmt.Sprintf("WRK [%d] KEY [%s]", n, key)
                m.Add(key, msg)
                k<- key
                operations++
            }
            fmt.Printf("WRK [%d] done [%d] operations...\n", n, operations)
            return
        }(&wg, keys, cmap, j)
    }
    wg.Wait()
    cmap.Range(func(k string, v []string) {
        fmt.Printf("For KEY [%s] found:\n", k)
        for idx, str := range v {
            fmt.Printf("\t> MSG [%d]: [%s]\n", idx + 1, str)
        }
        fmt.Println()
    })
}
