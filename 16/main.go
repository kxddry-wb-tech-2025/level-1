package main

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivotIdx := rand.Intn(len(arr)) // вместо середины или первого элемента выбираем рандомный
	pivot := arr[pivotIdx]
	var left, right []int
	for i, v := range arr {
		if i == pivotIdx {
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
