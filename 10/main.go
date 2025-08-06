package main

import "fmt"

func Group(seq []float64) map[int][]float64 {
	if len(seq) == 0 {
		return nil
	}
	mp := make(map[int][]float64)

	for _, temp := range seq {
		x := int(temp) - int(temp)%10
		mp[x] = append(mp[x], temp)
	}
	return mp
}

func main() {
	seq := []float64{-25.4, -27., 13., 19., 15.5, 24.5, -21.0, 32.5}
	mp := Group(seq)

	for k, v := range mp {
		fmt.Println(k, v)
	}
}
