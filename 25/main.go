package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	select {
	case <-time.After(duration):
		return
	}
}

const N = 3

func main() {
	sleep(N * time.Second)
	fmt.Println(N, "seconds passed")
}
