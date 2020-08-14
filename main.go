package main

import (
	"fmt"

	"github.com/joeunha/learn-go/syntax"
)

func main() {
	dict := syntax.Dictionary{}
	dict["name"] = "joeun"

	value, err := dict.Search("age")

	if err != nil {
		fmt.Println(err, ":", value)
	} else {
		fmt.Println(value)
	}
}
