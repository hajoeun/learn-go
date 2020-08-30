package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [5]string{"foo", "bar", "vvv", "kkk", "sss"}

	for _, person := range people {
		go isPerson(person, channel)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("wating for ", i)
		fmt.Println(<-channel) // blocking operation
	}
}

func isPerson(person string, channel chan string) {
	time.Sleep(time.Second * 2)
	channel <- person + "'s send from channel"
}
