package main

import "fmt"

func ListIntersections(l1, l2 []int) []int {
	freq := make(map[int]int)
	for _, v := range l1 {
		freq[v]++
	}
	var out []int
	for _, v := range l2 {
		if freq[v] > 0 {
			freq[v]--
			out = append(out, v)
		}
	}
	return out
}

func main() {
	ans := ListIntersections([]int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(ans)
}
