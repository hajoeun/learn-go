// Package headfirst is example of Head First Go.
package headfirst

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// GuessNumber is a game that guess a random of number.
// You can call this function with max number. Then you have to guess a random number between 1 and max that you call.
//   GuessNumber(100)
//   I've chosen a random number between 1 and 100
func GuessNumber(max int) {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(max) + 1
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("I've chosen a ramdom number between 1 and", max)
	fmt.Println("Can you guess it?")

	for i := 10; i > 0; i-- {
		fmt.Println("You have", i, "guesses left.")
		fmt.Print("Make a guess: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Oops, Your guess was HIGH.")
			fmt.Println("=========================")
		} else if guess < target {
			fmt.Println("Oops, Your guess was LOW.")
			fmt.Println("=========================")
		} else {
			fmt.Println("Good job! You guessed it!")
			return
		}
	}

	fmt.Println("Sorry, You didn't guess my number. It was:", target)
}
