package headfirst

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFile is example read file
func ReadFile() {
	file, err := os.Open("./headfirst/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
