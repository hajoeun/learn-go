package syntax

import "fmt"

// Slice function
func Slice() {
	array := [3]int{}
	array[0] = 1
	fmt.Println(array) // [1 0 0]

	for i := range array {
		array[i] = i + 10
	}
	fmt.Println(array) // [10 11 12]

	slice1 := []int{}
	// slice1[2] = 30 // runtime error: index out of range [2] with length 0
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2) // [1 2 3 4 5]

	for i, v := range slice2 {
		slice2[i] = v + 10
	}
	fmt.Println(slice2) // [11 12 13 14 15]
}
