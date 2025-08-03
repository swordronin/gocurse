package basics

import "fmt"

func main() {

	// simple iteration over a range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// iterate over collection
	// number := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range number {
	// 	fmt.Printf("Index: %d, Value %d\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue the loop but skip the rest of the lines/statements
	// 	}
	// 	fmt.Println("Odd number:", i)
	// 	if i == 5 {
	// 		break //break out of the loop
	// 	}
	// }

	// Asteriks Layout
	// rows := 5
	//Outer loop
	// for i := 1; i <= rows; i++ {
	// inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // moving to the next line
	// }
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off!")
}
