package main

import "fmt"

func sum1(a int, b int, c int) int {
	return a + b + c
}

func sum2(numbers [3]int) (total int) {
	number := 0
	for i := 0; i < len(numbers); i++ {
		number = numbers[i]
		total += number
	}
	return total
}

func sum3(numbers ...int) (total int) {
	number := 0
	for i := 0; i < len(numbers); i++ {
		number = numbers[i]
		total += number
	}
	return total
}

func sum4(numbers ...int) (total int) {
	for i := range numbers {
		total += numbers[i]
	}
	return total
}

func sum5(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(
		sum1(1, 2, 3),
		sum2([3]int{1, 2, 3}),
		sum3(1, 2, 3, 4, 5),
		sum4(1, 2, 3, 4, 5),
		sum5(1, 2, 3, 4, 5),
	)
}
