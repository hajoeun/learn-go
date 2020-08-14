package main

import (
	"fmt"

	"github.com/joeunha/learn-go/syntax"
)

func main() {
	dict := syntax.Dictionary{}
	dict["name"] = "joeun"
	fmt.Println(dict["name"])
}
