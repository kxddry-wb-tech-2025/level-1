package main

import "fmt"

func Generate(a []int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, x := range a {
			ch <- x
		}
		close(ch)
	}()

	return ch
}

func Double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			x, ok := <-in
			if !ok {
				return
			}
			out <- x * 2

		}
	}()
	return out
}

func Print(in <-chan int) <-chan struct{} {
	out := make(chan struct{})

	go func() {
		defer close(out)
		for {
			x, ok := <-in
			if !ok {
				return
			}
			fmt.Println(x)
		}
	}()

	return out
}

func main() {
	done := Print(Double(Generate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})))
	<-done
	fmt.Println("done")
}
