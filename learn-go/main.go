package main

import "fmt"

func main() {
	value := 10
	address := &value
	fmt.Println(value, address)

	fmt.Println(&value == address) // true
	fmt.Println(&value, address)

	// fmt.Println(&value == &address) // invalid operation: &value == &address (mismatched types *int and **int)
	fmt.Println(&value, &address)

	fmt.Println(value == *address) // true
	fmt.Println(value, *address)

	value = 20
	fmt.Println(value, *address)

	*address = 30
	fmt.Println(value, *address)

	// address = 40 // cannot use 40 (type int) as type *int in assignment
	// fmt.Println(value, *address)
}
