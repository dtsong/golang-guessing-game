package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Please input your guess")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Reading Input Error", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid Input, must be an integer", err)
	}

	fmt.Println("Your guess is", guess)

	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Try again.")
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Try again.")
	} else {
		fmt.Println("Correct, you Legend!")
	}
}
