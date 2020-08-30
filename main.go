package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	people := [2]string{"foo", "bar"}

	for _, person := range people {
		go isPerson(person, channel)
	}

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func isPerson(person string, channel chan bool) {

	if person == "foo" {
		time.Sleep(time.Second * 5)
		channel <- true
	} else {
		channel <- false
	}

}
