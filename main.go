package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hi, Welcome to the guessing game. Your goal is to guess an integer between 0-10")

	secret := getRandomNumber()

	for matching := false; !matching; {
		guess := getUserInput()
		matching = isMatching(secret, guess)
	}
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())

	return rand.Int() % 11
}

func getUserInput() int {
	var input int

	fmt.Println("Please input your guess: ")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input...")
		return input
	}

	fmt.Println("You guessed: ", input)

	return input
}

func isMatching(secret, guess int) bool {
	if guess == secret {
		fmt.Println("You Win!")
		return true
	}

	if guess > secret {
		fmt.Println("Too high")
		return false
	}

	fmt.Println("Too low")
	return false
}
