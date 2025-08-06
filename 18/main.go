package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	n atomic.Int64
}

func (c *Counter) Inc() {
	c.n.Add(1)
}

func (c *Counter) Dec() {
	c.n.Add(-1)
}

func (c *Counter) Val() int64 {
	return c.n.Load()
}

func main() {
	wg := sync.WaitGroup{}
	cnt := &Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				cnt.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println(cnt.Val())
}
