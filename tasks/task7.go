package main

import (
    "fmt"
    "sync"
)

// Тип для конкурентной работы с map
// Используем RWMutex так как
// планируем параллельно модифицировать и читать значения
type ConcurrentMap struct {
    lock sync.RWMutex
    coll map[string][]string
}

func (c *ConcurrentMap) Add(key, value string) {
    // Отложенный вызов разблокировки мьютекса
    defer c.lock.Unlock()
    // берем мьютекс, все прочие писатели / читатели ждут
    // пока мы освободим ресурс
    c.lock.Lock()
    if _, contains := c.coll[key]; contains == true {
        c.coll[key] = append(c.coll[key], value)
    } else {
        c.coll[key] = []string{value}
    }
    return
}

func (c *ConcurrentMap) Get(key string) []string {
    // Блокируем писателей, но можем читать
    defer c.lock.RUnlock()
    c.lock.RLock()
    return c.coll[key]
}

func (c *ConcurrentMap) Range(fn func(key string, val []string)) {
    // Можно передать функцию для вывода
    // всех значений
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
    // Используем WaitGroup для синхронизации
    var wg sync.WaitGroup
    var mapkeys []string = []string{"A", "B", "C", "D"}
    cmap := NewConcMap()
    wcount := 8
    // Создаем пул воркеров и заполняем токены
    keys := make(chan string, 4)
    for i := 0; i < 4; i++ {
        keys<- mapkeys[i]
    }
    for j := 0; j < wcount; j++ {
        wg.Add(1)
        // Запускаем 8 воркеров, но токена всего 4
        // поэтому остальные 4 ждут
        go func(w *sync.WaitGroup, k chan string, m *ConcurrentMap, n int) {
            defer w.Done()
            var operations int = 0
            for i := 0; i < 10; i++ {
                // берем токен
                key := <-k
                msg := fmt.Sprintf("WRK [%d] KEY [%s]", n, key)
                // пишем в map.
                // Пока один пишет остальные ждут
                m.Add(key, msg)
                // Возвращаем токен.
                // Это всегда безопасно (в плане блокировки)
                // так как невозможно вернуть больше, чем взяли.
                // Поскольку канал буферизованный, мы вновь наполним его
                // токенами
                k<- key
                operations++
            }
            fmt.Printf("WRK [%d] done [%d] operations...\n", n, operations)
            return
        }(&wg, keys, cmap, j)
    }
    // Дожидаемся, пока все 8 горутин отработают
    wg.Wait()
    cmap.Range(func(k string, v []string) {
        fmt.Printf("For KEY [%s] found:\n", k)
        for idx, str := range v {
            fmt.Printf("\t> MSG [%d]: [%s]\n", idx + 1, str)
        }
        fmt.Println()
    })
}
