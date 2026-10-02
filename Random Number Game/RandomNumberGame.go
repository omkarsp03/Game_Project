package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Random Number Game")
	scanner := bufio.NewScanner(os.Stdin)
	i := true
	for i != false {
		fmt.Println(" Enter a Guess ")
		scanner.Scan()
		input, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		n := rand.Intn(10)

		if input != int64(n) {
			fmt.Println("You Guessed Wrong")
		} else {
			fmt.Println("You Guessed Right")
			i = false
		}
	}
}
