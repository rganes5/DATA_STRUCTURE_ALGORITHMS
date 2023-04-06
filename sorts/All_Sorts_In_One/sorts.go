package main

import (
	"fmt"
)

func main() {
	var choice int
	array := []int{8, 3, 9, 1, 6, 7, 5, 2, 4}
	fmt.Println("Welcome to Sorting")
	fmt.Printf("The unsorted array is  :\n%v\n", array)
	fmt.Println("Enter the sort that you would like to perform:")
	fmt.Print("1 for BubbleSort\n2 for InsertionSort\n3 for selectionSort\n4 for QuickSort\n5 for MergeSort\n")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		bubblesort(array)
	case 2:
		InsertionSort(array)
	case 3:
		selectionSort(array)
	case 4:
		quickSort(array, 0, len(array)-1)
		fmt.Println(array)
	case 5:
		mergeSort(array)
		fmt.Println(array)
	default:
		fmt.Println("Enter correct option")
	}
}

//bubbleSort

func bubblesort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Printf("\nSorted array using Bubble Sort is: \n%v", array)
}

//InsertionSort

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
	fmt.Printf("\nSorted array using Insertion Sort is: \n%v", array)
}

//selectionSort

func selectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]

	}
	fmt.Printf("\nSorted array using Selection Sort is: \n%v", array)
}

//quickSort

func quickSort(array []int, start int, end int) {
	if end <= start {
		return
	}
	pivot := partition(array, start, end)
	quickSort(array, start, pivot-1)
	quickSort(array, pivot+1, end)
}
func partition(array []int, start int, end int) int {
	pivot := array[end]
	i := start - 1

	for j := start; j <= end; j++ {
		if array[j] < pivot {
			i++
			array[j], array[i] = array[i], array[j]
		}
	}
	i++
	array[i], array[end] = array[end], array[i]
	return i
}

// mergeSort
func mergeSort(array []int) {
	length := len(array)
	mid := length / 2
	if length <= 1 {
		return
	}

	leftArray, rightArray := []int{}, []int{}
	for i, v := range array {
		if i < mid {
			leftArray = append(leftArray, v)
		} else {
			rightArray = append(rightArray, v)
		}
	}
	mergeSort(leftArray)
	mergeSort(rightArray)
	merge(leftArray, rightArray, array)

}
func merge(leftArray []int, rightArray []int, array []int) {
	leftSize := len(array) / 2
	rightSize := len(array) - leftSize

	i, l, r := 0, 0, 0
	for l < leftSize && r < rightSize {
		if leftArray[l] < rightArray[r] {
			array[i] = leftArray[l]
			i++
			l++
		} else {
			array[i] = rightArray[r]
			i++
			r++
		}
	}
	for l < leftSize {
		array[i] = leftArray[l]
		i++
		l++

	}
	for r < rightSize {
		array[i] = rightArray[r]
		i++
		r++
	}

}
