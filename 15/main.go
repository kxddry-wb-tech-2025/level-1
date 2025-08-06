package main

var justString string

func createHugeString(n int) string {
	return "" // placeholder
}

// memory leak chance
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

// correct way
func correctFunc() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	for i := range 100 { // copy data into a temporary byte slice
		tmp[i] = v[i]
	}
	justString = string(tmp)
}

func main() {
	correctFunc()
}
