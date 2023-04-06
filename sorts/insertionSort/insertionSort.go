package main

import "fmt"

func main() {
	userSlices := []int{9, 2, 8, 3, 6, 1, 5, 4, 7}
	insertionSort(userSlices)

}
func insertionSort(userSlices []int) {
	for i := 1; i < len(userSlices); i++ {
		temp := userSlices[i]
		for j := 0; j < i; j++ {
			if temp < userSlices[j] {
				userSlices[i], userSlices[j] = userSlices[j], userSlices[i]
			}
		}
	}
	fmt.Println(userSlices)
}
