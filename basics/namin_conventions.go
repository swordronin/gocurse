package basics

import "fmt"

func main() {

	type Employee struct {
		FirstName string
		LastName  string
		Age       int
	}
	// Pascal case
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Struct, interfacer, enums
	//
	// Snake_case
	// EG. user_id, first_name, http_request
	// UPPERCASE
	// use case is Constants

	// mixedCase
	// Eg. javaScript, htmlDocuments, isValid
	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID:", employeeID)
}
