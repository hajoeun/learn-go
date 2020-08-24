package main

import (
	"fmt"
	"time"
)

func main() {
	go count("foo", 10)
	go count("bar", 5)
	count("uuu", 5)
}

func count(name string, len int) {
	for i := 0; i < len; i++ {
		fmt.Println(name, "count", i)
		time.Sleep(time.Second)
	}
}
