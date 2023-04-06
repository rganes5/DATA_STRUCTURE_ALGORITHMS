package main

import "fmt"

func main() {
	userSlices := []int{9, 2, 8, 3, 6, 1, 5, 4, 7}
	selectionSort(userSlices)

}

func selectionSort(userSlices []int) {
	for i := 0; i < len(userSlices)-1; i++ {
		min := i
		for j := i + 1; j < len(userSlices); j++ {
			if userSlices[j] < userSlices[min] {
				min = j
			}
		}
		temp := userSlices[i]
		userSlices[i] = userSlices[min]
		userSlices[min] = temp

	}
	fmt.Println(userSlices)
}
