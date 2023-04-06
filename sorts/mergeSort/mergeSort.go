package main

import "fmt"

func main() {
	fmt.Println("welcome to merge sort")
	userSlices := []int{8, 2, 4, 6, 9, 5, 1, 3, 7}
	fmt.Printf("\nThe original array is:%v\n", userSlices)
	mergeSort(userSlices)
	fmt.Printf("\nThe sorted array is%v", userSlices)
}

//SPLITTING

func mergeSort(userSlices []int) {
	mid := len(userSlices) / 2
	if len(userSlices) <= 1 {
		return
	}
	leftArray, rightArray := []int{}, []int{}
	for i, v := range userSlices {
		if i < mid {
			leftArray = append(leftArray, v)
		} else {
			rightArray = append(rightArray, v)
		}
	}

	mergeSort(leftArray)
	mergeSort(rightArray)
	merge(leftArray, rightArray, userSlices)
}

//MERGING

func merge(leftArray, rightArray, userSlices []int) {
	fmt.Println(userSlices)
	leftLength := len(userSlices) / 2 //leftLength:=len(leftArray)
	rightLength := len(userSlices) - leftLength
	r, l, i := 0, 0, 0

	for l < leftLength && r < rightLength {
		if leftArray[l] < rightArray[r] {
			userSlices[i] = leftArray[l]
			i++
			l++
		} else {
			userSlices[i] = rightArray[r]
			i++
			r++
		}
	}

	for l < leftLength {
		userSlices[i] = leftArray[l]
		i++
		l++
	}
	for r < rightLength {
		userSlices[i] = rightArray[r]
		i++
		r++
	}

}
