package main

func Swap1(x, y *int) {
	*x, *y = *y, *x
}

func Swap2(x, y *int) {
	*x ^= *y
	*y ^= *x
	*x ^= *y
}

func main() {}
