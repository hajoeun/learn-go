package syntax

import "fmt"

// Array is syntax example for array
func Array() {
	var numbers [5]int
	fmt.Printf("%#v\n", numbers)

	// array literal
	var notes [3]string = [3]string{"do", "re", "mi"}
	fmt.Printf("%#v\n", notes)

	// array literal + short variable declaration
	colors := [3]string{
		"red",
		"green",
		"blue",
	}
	fmt.Printf("%#v\n", colors)

	// loop - len
	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	// loop - for ...range
	for i, color := range colors {
		fmt.Println(i, color)
	}
}
