// Instructions
// Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.

// The square of the sum of the first ten natural numbers is (1 + 2 + ... + 10)² = 55² = 3025.

// The sum of the squares of the first ten natural numbers is 1² + 2² + ... + 10² = 385.

// Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares of the first ten natural numbers is 3025 - 385 = 2640.

// You are not expected to discover an efficient solution to this yourself from first principles; research is allowed, indeed, encouraged. Finding the best algorithm for the problem is a key skill in software engineering.

package diffsquares

import (
	"fmt"
)

func SquareOfSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		// fmt.Printf("%d\n", i)
		result += i
		//55 square is 3025
	}

	return result * result
}

func SumOfSquares(n int) int {
	resultstores := 0
	for i := 0; i <= n; i++ {
		// fmt.Println(" SumOfSquares", i)
		resultstores += i * i
		// fmt.Println("result", resultstores)

	}
	return resultstores
}

func Difference(n int) int {
	f := SquareOfSum(n)
	d := SumOfSquares(n)
	g := f - d

	return g
}

func main() {
	x := 10
	// var x int
	// fmt.Println("Enter value to sum of square !!")
	// fmt.Scanln(&x)
	SquareOfSumOutput := SquareOfSum(x)
	SumOfSquaresOutput := SumOfSquares(x)
	DifferenceValue := Difference(x)
	fmt.Printf("SquareOfSum_Output : %d\nSumOfSquaresOutput : %d\nDifference :%d\n", SquareOfSumOutput, SumOfSquaresOutput, DifferenceValue)

}
