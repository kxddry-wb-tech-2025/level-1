package main

import (
	"context"
	"fmt"
)

func Solution(ctx context.Context, n int, ch <-chan any) {
	for i := range n {
		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				val, ok := <-ch
				if !ok {
					return
				}
				fmt.Println("worker", i+1, "received", val)
			}
		}(i)
	}
}

func main() {
	ch := make(chan any)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	Solution(ctx, 10, ch)
	for i := range 20 {
		ch <- i
	}
	close(ch)
}
