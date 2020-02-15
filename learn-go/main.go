package main

import "fmt"

func isEven(n int) (bool, int) {
	if v := n % 2; v == 0 {
		return true, v
	} else {
		return false, v
	}
}

func isOdd(n int) (bool, int) {
	v := n % 2
	if v == 0 {
		return false, v
	}
	return true, v
}

func main() {
	number := 10
	isEvenNumber, evenResult := isEven(number)
	isOddNumber, oddResult := isOdd(number)

	fmt.Println(isEvenNumber, evenResult)
	fmt.Println(isOddNumber, oddResult)
}
