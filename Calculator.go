package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to CLI Calculator")
	fmt.Println("-------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter first number: ")
	scanner.Scan()
	num1, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Invalid first number")
		return
	}

	fmt.Print("Enter second number: ")
	scanner.Scan()
	num2, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Invalid second number")
		return
	}

	fmt.Print("Enter operation (+, -, *, /): ")
	scanner.Scan()
	operation := strings.TrimSpace(scanner.Text())

	switch operation {
	case "+":
		fmt.Println("Addition =", num1+num2)

	case "-":
		fmt.Println("Subtraction =", num1-num2)

	case "*":
		fmt.Println("Multiplication =", num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
			return
		}
		fmt.Println("Division =", num1/num2)

	default:
		fmt.Println("Invalid operation")
	}
}
