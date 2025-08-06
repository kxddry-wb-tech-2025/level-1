package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SafeMap struct {
	mp map[any]any
	mu sync.RWMutex
}

func NewMap() SafeMap {
	return SafeMap{
		mp: make(map[any]any),
		mu: sync.RWMutex{},
	}
}

func (m *SafeMap) Set(key any, value any) {
	m.mu.Lock()
	m.mp[key] = value
	m.mu.Unlock()
}

func (m *SafeMap) Get(key any) (any, bool) {
	m.mu.RLock()
	v, ok := m.mp[key]
	m.mu.RUnlock()
	return v, ok
}

// go run -race .
func main() {
	opt1 := sync.Map{}
	opt2 := NewMap()

	keys := []any{1, 3, 4, "hello", "world", "sss", 2}
	vals := []any{2, 4, 6, "test", "testov", "testovich", 9}

	wg := sync.WaitGroup{}

	for i := range keys {
		wg.Add(1)
		k := keys[i]
		v := vals[rand.Intn(len(vals))]
		go func() {
			defer wg.Done()
			opt1.Store(k, v)
			opt2.Set(k, v)
		}()
	}
	wg.Wait()

	for _, k := range keys {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a, _ := opt1.Load(k)
			b, _ := opt2.Get(k)

			fmt.Println(a, b)
		}()
	}
	wg.Wait()
}
