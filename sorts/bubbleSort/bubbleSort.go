package main

import "fmt"

func main() {
	var key int
	userSlices := []int{9, 2, 4, 6, 1, 7, 5, 3, 8}
	bubbleSort(userSlices, key)
}

func bubbleSort(userSlices []int, key int) {
	fmt.Println(len(userSlices))
	for i := 0; i < len(userSlices)-1; i++ {
		flag := 0
		for j := 0; j < len(userSlices)-i-1; j++ {
			if userSlices[j] > userSlices[j+1] {
				userSlices[j], userSlices[j+1] = userSlices[j+1], userSlices[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	fmt.Println("Sorted: ")
	fmt.Println(userSlices)
}
