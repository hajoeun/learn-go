package headfirst

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GradeChecker is check a grade
func GradeChecker() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade > 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
