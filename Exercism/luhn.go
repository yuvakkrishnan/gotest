package main

import (
	"fmt"
	"strings"
)

func isValid(number string) bool {
	// Strip spaces from the input
	number = strings.Replace(number, " ", "", -1)
	// Check if the input is too short
	if len(number) <= 1 {
		return false
	}
	// Check for any non-digit characters
	for _, c := range number {
		if c < '0' || c > '9' {
			return false
		}
	}
	// Perform the Luhn algorithm
	sum := 0
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if i%2 == len(number)%2 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	// Check if the sum is divisible by 10
	return sum%10 == 0
}

func main() {
	fmt.Println(isValid("4539 3195 0343 6467")) // true
	fmt.Println(isValid("8273 1232 7352 0569")) // false
}

// Instructions
// Given a number determine whether or not it is valid per the Luhn formula.

// The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.

// The task is to check if a given string is valid.

// Validating a Number
// Strings of length 1 or less are not valid. Spaces are allowed in the input, but they should be stripped before checking. All other non-digit characters are disallowed.

// Example 1: valid credit card number
// 4539 3195 0343 6467
// The first step of the Luhn algorithm is to double every second digit, starting from the right. We will be doubling

// 4_3_ 3_9_ 0_4_ 6_6_
// If doubling the number results in a number greater than 9 then subtract 9 from the product. The results of our doubling:

// 8569 6195 0383 3437
// Then sum all of the digits:

// 8+5+6+9+6+1+9+5+0+3+8+3+3+4+3+7 = 80
// If the sum is evenly divisible by 10, then the number is valid. This number is valid!

// Example 2: invalid credit card number
// 8273 1232 7352 0569
// Double the second digits, starting from the right

// 7253 2262 5312 0539
// Sum the digits

// 7+2+5+3+2+2+6+2+5+3+1+2+0+5+3+9 = 57
// 57 is not evenly divisible by 10, so this number is not valid.

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func luhnAlgorithm(number string) bool {
// 	// Step 1: Convert the number to slice of integers
// 	digits := make([]int, len(number))
// 	// fmt.Println("Len of Num", len(number))
// 	for i := range digits {
// 		digits[i] = int(number[i] - '0')
// 		//The ASCII value of the character '0' is subtracted from the ASCII value of the current character. Since the ASCII value of '0' is 48, and the ASCII values of the digits '0' to '9' are in consecutive order starting from 48, this subtraction will give the integer value of the digit represented by the current character.
// 		// fmt.Println("Digits", digits[i])

// 		fmt.Println("Number ", number[i])
// 	}
// 	// fmt.Println("Len of digits", len(digits))
// 	// Step 2: Double every second digit
// 	for i := len(digits) - 2; i >= 0; i -= 2 {
// 		// fmt.Println("value of i", i)
// 		// fmt.Println("digits of i 1st", digits[i])
// 		digits[i] *= 2
// 		// fmt.Println("digits of I", digits[i])

// 		if digits[i] > 9 {
// 			digits[i] -= 9
// 			// fmt.Println("after minus Check", digits)

// 		}
// 	}
// 	// fmt.Println("Before Check", digits)
// 	// Step 3: Check if the sum is divisible by 10
// 	sum := 0
// 	for _, digit := range digits {
// 		sum += digit
// 		// fmt.Println("Sum", sum)
// 	}
// 	return sum%10 == 0
// }

// func main() {
// 	number := "4539319503436467"
// 	valid := luhnAlgorithm(number)
// 	fmt.Println("Number:", number)
// 	fmt.Println("Valid: ", strconv.FormatBool(valid))
// }
