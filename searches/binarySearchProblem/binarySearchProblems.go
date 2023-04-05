package main

import "fmt"

func main() {
	userSlices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(binarySearch(userSlices, 5))
}

func binarySearch(userSlices []int, key int) bool {
	low := 0
	high := len(userSlices) - 1
	for low <= high {
		mid := (low + high) / 2
		if key == userSlices[mid] {
			return true
		} else if key > userSlices[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
