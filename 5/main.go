package main

import (
	"context"
	"fmt"
	"time"
)

func Producer(ctx context.Context, value any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- value:
			}
		}
	}()
	return ch
}

func Consumer(ctx context.Context, ch <-chan any) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			_ = v // do something with the value
		}
	}
}

const N = 3

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()
	go Consumer(ctx, Producer(context.Background(), "hello world"))
	<-ctx.Done()
	fmt.Println(N, "seconds passed -> done")

}
