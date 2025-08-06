package main

import (
	"fmt"
	"sync"
)

func Calculation(n int) int {
	return n * n
}

func Solution1(arr []int) {
	wg := sync.WaitGroup{}
	for _, a := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqr := Calculation(a)
			fmt.Println(sqr)
		}()
	}
	wg.Wait()
}

func Solution2(arr []int) {
	wg := sync.WaitGroup{}
	out := make([]int, len(arr))
	for i, a := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i] = Calculation(a)
		}()
	}
	wg.Wait()

	for _, a := range out {
		fmt.Println(a)
	}
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	// Solution1(arr[:]) -- может быть другой порядок
	Solution2(arr[:])
}
