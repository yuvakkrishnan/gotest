package main

import (
	"fmt"
	"strconv"
)

func luhnAlgorithm(number string) bool {
	// Step 1: Convert the number to slice of integers
	digits := make([]int, len(number))
	fmt.Println("Len of Num", len(number))
	for i := range digits {
		digits[i] = int(number[i] - '0')
		//The ASCII value of the character '0' is subtracted from the ASCII value of the current character. Since the ASCII value of '0' is 48, and the ASCII values of the digits '0' to '9' are in consecutive order starting from 48, this subtraction will give the integer value of the digit represented by the current character.
		// fmt.Println("Digits", digits[i])

		// fmt.Println("Number ", number[i])
	}
	fmt.Println("Len of digits", len(digits))
	// Step 2: Double every second digit
	for i := len(digits) - 2; i >= 0; i -= 2 {
		fmt.Println("value of i", i)
		fmt.Println("digits of i 1st", digits[i])
		digits[i] *= 2
		fmt.Println("digits of I", digits[i])

		if digits[i] > 9 {
			digits[i] -= 9
			fmt.Println("after minus Check", digits)

		}
	}
	// fmt.Println("Before Check", digits)
	// Step 3: Check if the sum is divisible by 10
	sum := 0
	for _, digit := range digits {
		sum += digit
		// fmt.Println("Sum", sum)
	}
	return sum%10 == 0
}

func main() {
	number := "49927398716"
	valid := luhnAlgorithm(number)
	fmt.Println("Number:", number)
	fmt.Println("Valid: ", strconv.FormatBool(valid))
}
