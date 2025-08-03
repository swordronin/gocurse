package main

import "fmt"

func main() {
	// Variables declaration

	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Reminder:", result)

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	// Overflow with signed Intergers
	var maxInt int64 = 9223372036854775807 // Max Value for Int64
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

}
