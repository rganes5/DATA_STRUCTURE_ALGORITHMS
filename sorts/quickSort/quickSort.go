package main

import "fmt"

func main() {
	userSlices := []int{8, 2, 4, 6, 9, 5, 1, 3, 7}
	fmt.Println("QuickSort")
	fmt.Println("The array before sorting is: ", userSlices)
	quickSort(userSlices, 0, len(userSlices)-1)
	fmt.Printf("\nThe sorted array is: %v\n", userSlices)
}

func quickSort(userSlices []int, start int, end int) {
	if end <= start { //basecase
		return
	}
	pivot := partition(userSlices, start, end)
	quickSort(userSlices, start, pivot-1)
	quickSort(userSlices, pivot+1, end)
}

func partition(userSlices []int, start int, end int) int {
	pivot := userSlices[end]
	i := start - 1
	for j := start; j <= end; j++ {
		if userSlices[j] < pivot {
			i++
			userSlices[j], userSlices[i] = userSlices[i], userSlices[j]
			// fmt.Println(j, i)
		}
	}
	i++
	userSlices[i], userSlices[end] = userSlices[end], userSlices[i]
	// fmt.Println(userSlices)
	// fmt.Println(pivot)
	return i
}
