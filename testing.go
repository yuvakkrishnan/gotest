package main

import (
	"fmt"
	"strconv"
)

func luhnAlgorithm(number string) bool {
	digits := make([]int, len(number))
	for i := range digits {
		digits[i] = int(number[i] - '0')
	}
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] *= 2
		if digits[i] > 9 {
			digits[i] -= 9
		}
	}
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	return sum%10 == 0
}
func main() {
	number := "4539319503436457"
	valid := luhnAlgorithm(number)
	fmt.Println("Number:", number)
	fmt.Println("Valid: ", strconv.FormatBool(valid))
}
