package main

import "fmt"

func isEven(n int) (bool, int) {
	v := n % 2
	switch {
	case v == 0:
		return true, v
	default:
		return false, v
	}
}

func isOdd(n int) (bool, int) {
	switch v := n % 2; v {
	case 0:
		return false, v
	default:
		return true, v
	}
}

func main() {
	number := 10
	isEvenNumber, evenResult := isEven(number)
	isOddNumber, oddResult := isOdd(number)

	fmt.Println(isEvenNumber, evenResult)
	fmt.Println(isOddNumber, oddResult)
}
