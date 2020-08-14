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

	err2 := dict.Update(word, "joeun ha")

	if err2 != nil {
		fmt.Println(err2, ":", word)
	} else {
		result2, _ := dict.Search(word)
		fmt.Println(result2)
	}

	dict.Delete(word)
	result, err3 := dict.Search(word)

	if err3 != nil {
		fmt.Println(err3, ":", word)
	}
}
