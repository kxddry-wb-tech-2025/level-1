package main

func BinarySearch(arr []int, val int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (r + l) / 2
		if arr[mid] == val {
			return mid
		}
		if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
}
