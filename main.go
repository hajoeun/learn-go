package main

import (
	"fmt"

	"github.com/joeunha/learn-go/syntax"
)

func main() {
	dict := syntax.Dictionary{}
	word := "name"
	definition := "joeun"
	err := dict.Add(word, definition)

	if err != nil {
		fmt.Println(err)
	}

	result, _ := dict.Search(word)
	fmt.Println(result)

	err2 := dict.Add(word, "joeun2")

	if err2 != nil {
		fmt.Println(err2, ":", word)
	}
}
