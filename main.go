package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min, max := 1, 100
	secretNumber := rand.Intn(max-min) + min
	fmt.Println("The secret number is", secretNumber)
}
