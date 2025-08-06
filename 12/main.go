package main

import "fmt"

func CreateSet(a []string) []string {
	mp := make(map[string]struct{})
	var out []string
	for _, v := range a {
		if _, ok := mp[v]; !ok {
			out = append(out, v)
		}
		mp[v] = struct{}{}
	}
	return out
}

func main() {
	set := CreateSet([]string{"cat", "dog", "cat", "cat", "tree"})
	fmt.Println(set)
}
