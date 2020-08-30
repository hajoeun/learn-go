package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	people := [5]string{"foo", "bar", "vvv", "kkk", "sss"}

	for _, person := range people {
		go isPerson(person, channel)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func isPerson(person string, channel chan string) {
	channel <- person + "'s send from channel"
}
