package main

import "fmt"

func main() {
	/* most of the main data types
	1. string
	2. bool
	3. int
		a. int int8 int16 int32 int64
		b. uint uint8 uint16 uint32 uint64 uintptr
	4. byte - same as uint8
	5. rune - same as uint32
	6. float32 float64
	7. complex64 complex128
	*/

	// using var keyword to create a variable
	var name string = "Mohamed"
	var age = 34
	// using constants using const keyword
	const height = 186.3

	fmt.Println("name: ", name, "\nage: ", age, "\nheight: ", height)
	// %T is used to print the type of a variable
	fmt.Printf("variable name is of %T type \n", name)
	fmt.Printf("variable age is of %T type \n", age)
	fmt.Printf("variable height is of %T type \n", height)
}
