package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func StopOnCondition() {
	done := false

	go func() {
		for !done {

		}
		fmt.Println("done")
	}()

	time.Sleep(time.Second)
	done = true
	time.Sleep(time.Second)

}

func StopWithChannel() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				goto after
			default:
			}
		}
	after:
		fmt.Println("done")
	}()

	time.Sleep(time.Second)
	close(done)
	time.Sleep(time.Second)

}

func StopWithContext(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}

func Goexit() {
	go func() {
		for i := range 100 {
			if i == 15 {
				runtime.Goexit()
				// unreachable
			}
		}
	}()
}

func Timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timeout done")
				return
			default:
			}
		}
	}()

	// cancels before the 2-second timestamp
	time.Sleep(2 * time.Second)
}

// Stop via main
func main() {
	go func() {
		// might end before this
		fmt.Println("started")
	}()
	// program quits
}
