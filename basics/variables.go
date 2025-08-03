package basics

import "fmt"

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"
	//middleName := "Mayor"
	fmt.Println(middleName)

	// count := 10
	// lastName := "Smith"

	// Default Values
	// Numeric Types : 0
	// Boolean Types : false
	// String Types : " "
	// Pointer, slices, maps function and structs:nil

	// ---- Scope

	//fmt.Println(printName())
	printName()

}

func printName() {
	firstName := "Micheal"
	fmt.Println(firstName)
}
