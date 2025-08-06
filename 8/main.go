package main

func SetBit(n int, bit uint8, val bool) int {
	if bit > 32 {
		panic("bit out of range")
	}
	if !val {
		n &= ^(1 << bit)
	} else {
		n |= (1 << bit)
	}
	return n
}

func main() {
	// false means 0
	SetBit(5, 1., false)
}
