package main

import (
	"fmt"
)

func main() {
	var number int
	userSlices := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Enter the element to be searched. ")
	fmt.Scan(&number)
	fmt.Println(linearSearch(userSlices, number))

}

func linearSearch(userSlices []int, number int) bool {
	for _, v := range userSlices {
		if number == v {
			return true
		}
	}
	return false
}
