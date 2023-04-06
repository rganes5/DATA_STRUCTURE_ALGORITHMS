package main

import (
	"fmt"
)

// POWER OF A NUMBER
func CalculatePower(num int, power int) int {
	if power < 1 {
		return 1
	}
	return num * CalculatePower(num, power-1)
}

// FACTORIAL
func factorial(num int) int {
	if num < 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	var num, power int

	fmt.Printf("Enter value of num: ")
	fmt.Scan(&num)
	fmt.Printf("Enter value of power: ")
	fmt.Scan(&power)

	fmt.Printf("%d to the power of %d is:\n", num, power)
	fmt.Println(CalculatePower(num, power))
}
