package syntax

import "fmt"

// Pointer should print pointer example
func Pointer() {
	var myInt int
	var myIntPointer *int // * for pointer type

	myInt = 42
	myIntPointer = &myInt // & for myInt pointer (address value)

	fmt.Println(*myIntPointer) // * for value at pointer
}
